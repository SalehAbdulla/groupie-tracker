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
	port    string
	mux     *http.ServeMux
	view    []constants.ArtistView
	httpc   *http.Client
	Artists []constants.ArtistData
	relIdx  *constants.RelationIndex
}

func New(port string) (*App, error) {
	a := &App{
		port: port,
		mux:  http.NewServeMux(),
		httpc: &http.Client{
			Timeout: 12 * time.Second,
		},
		Artists: nil,
		relIdx:  nil,
	}

	var wg sync.WaitGroup
	errs := make(chan error, 2)

	wg.Add(2)

	go func() {
		defer wg.Done()
		errs <- fetchJSON(a.httpc, "https://groupietrackers.herokuapp.com/api/artists", &a.Artists)
	}()

	go func() {
		defer wg.Done()
		errs <- fetchJSON(a.httpc, "https://groupietrackers.herokuapp.com/api/relation", &a.relIdx)
	}()

	wg.Wait()
	close(errs)

	//fmt.Println(a.Artists)
	//fmt.Println(a.relIdx)

	for e := range errs {
		if e != nil {
			return nil, e
		}
	}

	relByID := make(map[int]map[string][]string, len(a.relIdx.Index))
	for _, x := range a.relIdx.Index {
		relByID[x.ID] = x.DatesLocations
	}

	view := make([]constants.ArtistView, 0, len(a.Artists))
	for _, ar := range a.Artists {
		view = append(view, constants.ArtistView{
			ArtistData: ar,
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
		http.StripPrefix("/templates/", http.FileServer(http.FS(h.Static))))
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
