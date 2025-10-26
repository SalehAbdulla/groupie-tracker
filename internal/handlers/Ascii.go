package handlers

import (
	"net/http"
)

func (h *Handlers) AsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// We’re returning HTML for all the template-rendered responses.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	h.render(w, "home.html", PageData{})
}
