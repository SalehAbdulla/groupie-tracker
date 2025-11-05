package handlers

import (
	"net/http"

	"groupie-tracker/internal/constants"
)

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		h.render(w, "ErrorPage.html", constants.Error{Error: http.StatusText(http.StatusMethodNotAllowed)})
		return
	}
	h.render(w, "home.html", constants.HomePageData{Artists: h.Artists})
}

func (h *Handlers) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	h.render(w, "ErrorPage.html", constants.Error{Error: http.StatusText(http.StatusNotFound)})
}
