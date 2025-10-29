package app

import (
	"encoding/json"
	"groupie-tracker/internal/constants"
	"groupie-tracker/internal/handlers"
	"net/http"
)

type App struct {
	port          string
	mux           *http.ServeMux
	data          []constants.ArtistData
}

func New(port string) (*App, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []constants.ArtistData
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, err
	}

	app := &App{
		port: port,
		mux:  http.NewServeMux(),
		data: artists,
	}

	app.routes()
	return app, nil
}

func (app *App) GetPort() string                 { return app.port }
func (app *App) GetData() []constants.ArtistData { return app.data }

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, pattern := app.mux.Handler(r)
	if pattern != "" {
		app.mux.ServeHTTP(w, r)
		return
	}
	h := handlers.New()
	w.WriteHeader(http.StatusNotFound)
	h.NotFound(w, r)
}

func (app *App) routes() {
	h := handlers.New()

	app.mux.HandleFunc("/", h.Home)
	app.mux.Handle("/templates/",
		http.StripPrefix("/templates/",
			http.FileServer(http.FS(h.Static()))))
}
