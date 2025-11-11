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
	h.Render(w, r, "home.html", constants.HomePageData{Artists: h.Artists})
}

func (h *Handlers) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	h.Render(w, r, "ErrorPage.html", constants.Error{Error: http.StatusText(http.StatusNotFound)})
}

func (h *Handlers) BadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	h.Render(w, r, "ErrorPage.html", constants.Error{Error: http.StatusText(http.StatusBadRequest)})
}

func (h *Handlers) InternalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	h.Render(w, r, "ErrorPage.html", constants.Error{Error: http.StatusText(http.StatusInternalServerError)})
}

func (h *Handlers) MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	h.Render(w, r, "ErrorPage.html", constants.Error{Error: http.StatusText(http.StatusMethodNotAllowed)})
}
