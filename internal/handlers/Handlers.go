package handlers

import (
	"html/template"
	"io/fs"
	"net/http"
	"strings"

	"groupie-tracker/internal/constants"
	"groupie-tracker/ui"
)

type Handlers struct {
	tpl     *template.Template
	static  fs.FS
	Artists []constants.ArtistView // Exported + matches how we’ll use it
}

func New(view []constants.ArtistView) *Handlers {
	t := template.Must(template.ParseFS(ui.Files, "templates/*.html"))
	sub, _ := fs.Sub(ui.Files, "templates")
	return &Handlers{tpl: t, static: sub, Artists: view}
}

func (h *Handlers) render(w http.ResponseWriter, name string, data any) {
	if err := h.tpl.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handlers) renderStr(w http.ResponseWriter, name string, data any) {
	const inlinePrefix = "inline:"
	if strings.HasPrefix(name, inlinePrefix) {
		src := strings.TrimPrefix(name, inlinePrefix)

		t, err := h.tpl.Clone()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		const inlineName = "__inline__"
		if _, err := t.New(inlineName).Parse(src); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := t.ExecuteTemplate(w, inlineName, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := h.tpl.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handlers) Static() fs.FS { return h.static }
