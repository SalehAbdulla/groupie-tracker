package handlers

import (
	"net/http"

	"groupie-tracker/internal/constants"
)

type HomePageData struct {
	Artists []constants.ArtistView
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.NotFound(w, r); return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed); return
	}
	h.render(w, "home.html", HomePageData{Artists: h.Artists})
}

func (h *Handlers) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	h.render(w, "ErrorPage.html", constants.ArtistData{})
}
