package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"groupie-tracker/internal/constants"
	"groupie-tracker/internal/handlers"
	"net/http"
	"sync"
	"time"
)

type App struct {
	port string
	mux  *http.ServeMux
	view []constants.ArtistView
}

func New(port string) (*App, error) {
	a := &App{
		port: port,
		mux:  http.NewServeMux(),
	}

	client := &http.Client{Timeout: 10 * time.Second}

	var artists []constants.ArtistData
	var relations constants.CardPageData
	var wg sync.WaitGroup
	var errA, errR error

	wg.Add(2)

	go func() {
		defer wg.Done()
		errA = fetchJSON(client, "https://groupietrackers.herokuapp.com/api/artists", &artists)
	}()

	go func() {
		defer wg.Done()
		errR = fetchJSON(client, "https://groupietrackers.herokuapp.com/api/relation", &relations)
	}()

	wg.Wait()
	if errA != nil || errR != nil {
		return nil, fmt.Errorf("fetch failed: %v %v", errA, errR)
	}

	// Merge artist + relation data
	relationMap := make(map[int]map[string][]string)
	for _, rel := range relations.Index {
		relationMap[rel.ID] = rel.DatesLocations
	}

	views := make([]constants.ArtistView, len(artists))
	for i, artist := range artists {
		views[i] = constants.ArtistView{
			ArtistData: artist,
			Rel:        relationMap[artist.ID],
		}
	}
	a.view = views

	a.routes()
	return a, nil
}

func (a *App) GetPort() string { return a.port }

func (a *App) routes() {
	h, err  := handlers.New(a.view)
	if err != nil {a.mux.HandleFunc("/", h.InternalServerError)}
	a.mux.HandleFunc("/", h.Home)
	a.mux.HandleFunc("/card-data", h.CardData)
	a.mux.Handle("/templates/",
		http.StripPrefix("/templates/", http.FileServer(http.FS(h.Static))))
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
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
