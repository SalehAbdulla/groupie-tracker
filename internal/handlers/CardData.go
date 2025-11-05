package handlers

import (
	"groupie-tracker/internal/constants"
	"strconv"

	"net/http"
)

func (h *Handlers) CardData(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if !(id >= 1 && id <= len(h.Artists)) || err != nil {
		h.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		h.render(w, "ErrorPage.html", constants.Error{Error: http.StatusText(http.StatusMethodNotAllowed)})
		return
	}
	h.render(w, "CardPageData.html", constants.HomePageData{Artists: h.Artists})
}
