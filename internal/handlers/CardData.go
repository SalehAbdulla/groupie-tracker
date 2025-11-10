package handlers

import (
	"net/http"
	"strconv"
)

func (h *Handlers) CardData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.MethodNotAllowed(w, r)
		return
	}
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 || id > len(h.Artists) {
		h.NotFound(w, r)
		return
	}
	h.render(w, "CardData.html", h.Artists[id-1])
}
