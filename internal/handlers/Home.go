package handlers

import "net/http"

type PageData struct {
	Id  int
	Image   string
	Name string
	Members  []string
	CreationDate int
	FirstAlbum string
	Locations string
	ConcertDates string
	Relations string
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {h.NotFound(w, r); return}
	if r.Method != http.MethodGet {http.Error(w, "method not allowed", http.StatusMethodNotAllowed); return}
	h.render(w, "home.html", PageData{})
}
