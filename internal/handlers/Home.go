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
		h.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	//h.render(w, "home.html", HomePageData{Artists: h.Artists})
	h.renderStr(w, "inline:<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n  <meta charset=\"UTF-8\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n  <title>Document</title>\n</head>\n<body>\n  <section class=\"grid\">\n  {{range .Artists}}\n  <article class=\"card\">\n    <img class=\"card-img\" src=\"{{.Image}}\" alt=\"{{.Name}}\">\n    <div class=\"card-body\">\n      <h2 class=\"card-title\">{{.Name}}</h2>\n      <p class=\"muted\">Since {{.CreationDate}}</p>\n      <p><strong>First album:</strong> {{.FirstAlbum}}</p>\n      {{if .Members}}\n      <div class=\"members\">\n        <strong>Members:</strong>\n        <ul>{{range .Members}}<li>{{.}}</li>{{end}}</ul>\n      </div>\n      {{end}}\n    </div>\n  </article>\n  {{end}}\n</section>\n</body>\n</html>", HomePageData{Artists: h.Artists})
}

func (h *Handlers) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	h.render(w, "ErrorPage.html", constants.ArtistData{})
}
