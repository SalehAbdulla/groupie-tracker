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
	port         string
	mux          *http.ServeMux
	view         []constants.ArtistView
	httpc        *http.Client
	HomePageData []constants.ArtistData
	CardPageData *constants.CardPageData
}

func New(port string) (*App, error) {
	a := &App{
		port: port,
		mux:  http.NewServeMux(),
		httpc: &http.Client{
			Timeout: 12 * time.Second,
		},
		HomePageData: nil,
		CardPageData: nil,
	}

	var wg sync.WaitGroup
	errs := make(chan error, 2)

	wg.Add(2)

	go func() {
		defer wg.Done()
		errs <- fetchJSON(a.httpc, "https://groupietrackers.herokuapp.com/api/artists", &a.HomePageData)
	}()

	go func() {
		defer wg.Done()
		errs <- fetchJSON(a.httpc, "https://groupietrackers.herokuapp.com/api/relation", &a.CardPageData)
	}()

	wg.Wait()
	close(errs)

	for e := range errs {
		if e != nil {
			return nil, e
		}
	}

	cardPageData := make(map[int]map[string][]string, len(a.CardPageData.PageData))
	for _, data := range a.CardPageData.PageData {
		cardPageData[data.ID] = data.DatesLocations
	}

	homePageDate := make([]constants.ArtistView, 0, len(a.HomePageData))
	for _, ar := range a.HomePageData {
		homePageDate = append(homePageDate, constants.ArtistView{
			ArtistData: ar,
			Rel:        cardPageData[ar.ID],
		})
	}

	a.view = homePageDate
	a.routes()
	return a, nil
}

func (a *App) GetPort() string { return a.port }

func (a *App) routes() {
	h := handlers.New(a.view)
	a.mux.HandleFunc("/", h.Home)
	a.mux.HandleFunc("/card-data", h.CardData)
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
