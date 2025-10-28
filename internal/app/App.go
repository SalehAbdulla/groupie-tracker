package app

import (
	"groupie-tracker/internal/constants"
	"groupie-tracker/internal/handlers"
	"net/http"
)

type App struct {
	port          string
	templatesPath string
	mux           *http.ServeMux
	data          constants.ArtistData
}

func New(port string) (*App, error) {
	FetchedData, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {return nil, err}

	app := &App{
		port:          port,
		templatesPath: constants.TEMPLATES_PATH,
		mux:           http.NewServeMux(),
		data: FetchedData,
	}

	app.routes()
	return app, nil
}

func (app *App) GetPort() string { return app.port }

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, pattern := app.mux.Handler(r)
	if pattern != "" {
		app.mux.ServeHTTP(w, r)
		return
	}
	h := handlers.New(app.templatesPath)
	w.WriteHeader(http.StatusNotFound)
	h.NotFound(w, r)
}

func (app *App) routes() {
	h := handlers.New(app.templatesPath)

	app.mux.HandleFunc("/", h.Home)

	fileServer := http.StripPrefix("/templates/", http.FileServer(http.Dir(app.templatesPath)))

	app.mux.HandleFunc("/templates/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/templates/" {
			http.NotFound(w, r)
			return
		}
		fileServer.ServeHTTP(w, r)
	})
}
