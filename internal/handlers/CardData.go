package handlers

import (
	"fmt"
	"groupie-tracker/internal/constants"

	"net/http"
)

func (h *Handlers) CardData(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != fmt.Sprintf("/CardPageData/{%d}", id) {
		h.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	h.render(w, "CardPageData.html", constants.CardPageData{})
}
