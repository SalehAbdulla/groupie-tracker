package handlers

import (
	"groupie-tracker/internal/constants"
	"net/http"
	"strconv"
)

func (h *Handlers) CardData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		h.render(w, "ErrorPage.html", constants.Error{Error: http.StatusText(http.StatusMethodNotAllowed)})
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		h.NotFound(w, r)
		return
	}

	if id > len(h.Artists) || h.Artists[id-1].ID != id {
		h.NotFound(w, r)
		return
	}

	artist := h.Artists[id-1]
	h.render(w, "CardData.html", artist)
}
