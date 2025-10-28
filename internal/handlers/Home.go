package handlers

import (
	"groupie-tracker/internal/constants"
	"net/http"
)

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {h.NotFound(w, r); return}
	if r.Method != http.MethodGet {http.Error(w, "method not allowed", http.StatusMethodNotAllowed); return}
	h.render(w, "home.html", constants.ArtistData{})
}

func (h *Handlers) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	h.render(w, "ErrorPage.html", constants.ArtistData{})
}
