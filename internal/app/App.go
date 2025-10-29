package app

import (
	"encoding/json"
	"errors"
	"groupie-tracker/internal/constants"
	"groupie-tracker/internal/handlers"
	"net/http"
	"sync"
	"time"
)

type App struct {
	port  string
	mux   *http.ServeMux
	view  []constants.ArtistView
	httpc *http.Client
}

func New(port string) (*App, error) {
	a := &App{
		port: port,
		mux:  http.NewServeMux(),
		httpc: &http.Client{
			Timeout: 12 * time.Second,
		},
	}

	// 1) Fetch all endpoints concurrently
	var (
		artists []constants.ArtistData
		locIdx  constants.LocationsIndex
		dateIdx constants.DatesIndex
		relIdx  constants.RelationIndex
	)

	var wg sync.WaitGroup
	errs := make(chan error, 4)

	wg.Add(4)
	go func() { defer wg.Done(); errs <- fetchJSON(a.httpc, "https://groupietrackers.herokuapp.com/api/artists", &artists) }()
	go func() { defer wg.Done(); errs <- fetchJSON(a.httpc, "https://groupietrackers.herokuapp.com/api/locations", &locIdx) }()
	go func() { defer wg.Done(); errs <- fetchJSON(a.httpc, "https://groupietrackers.herokuapp.com/api/dates", &dateIdx) }()
	go func() { defer wg.Done(); errs <- fetchJSON(a.httpc, "https://groupietrackers.herokuapp.com/api/relation", &relIdx) }()
	wg.Wait()
	close(errs)

	for e := range errs {
		if e != nil {
			return nil, e
		}
	}

	// 2) Index by id for quick merge
	locByID := make(map[int][]string, len(locIdx.Index))
	for _, x := range locIdx.Index {
		locByID[x.ID] = x.Locations
	}
	dtByID := make(map[int][]string, len(dateIdx.Index))
	for _, x := range dateIdx.Index {
		dtByID[x.ID] = x.Dates
	}
	relByID := make(map[int]map[string][]string, len(relIdx.Index))
	for _, x := range relIdx.Index {
		relByID[x.ID] = x.DatesLocations
	}

	// 3) Merge into view slice (keep artists’ order)
	view := make([]constants.ArtistView, 0, len(artists))
	for _, ar := range artists {
		view = append(view, constants.ArtistView{
			ArtistData: ar,
			Locs:       locByID[ar.ID],
			Dts:        dtByID[ar.ID],
			Rel:        relByID[ar.ID],
		})
	}

	a.view = view
	a.routes()
	return a, nil
}

func (a *App) GetPort() string { return a.port }

func (a *App) routes() {
	h := handlers.New(a.view) 
	a.mux.HandleFunc("/", h.Home)
	a.mux.Handle("/templates/",
		http.StripPrefix("/templates/", http.FileServer(http.FS(h.Static()))))
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, pattern := a.mux.Handler(r); pattern != "" {
		a.mux.ServeHTTP(w, r)
		return
	}
	h := handlers.New(a.view)
	h.NotFound(w, r)
}

func fetchJSON[T any](c *http.Client, url string, dst *T) error {
	resp, err := c.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	return json.NewDecoder(resp.Body).Decode(dst)
}
