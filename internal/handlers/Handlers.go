package handlers

import (
	"html/template"
	"io/fs"
	"net/http"

	"groupie-tracker/ui"
)

type Handlers struct {
	tpl    *template.Template
	static fs.FS
}

func New() *Handlers {
	t := template.Must(template.ParseFS(ui.Files, "templates/*.html"))
	sub, _ := fs.Sub(ui.Files, "templates")
	return &Handlers{tpl: t, static: sub}
}

func (h *Handlers) render(w http.ResponseWriter, name string, data any) {
	if err := h.tpl.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handlers) Static() fs.FS {
	return h.static
}
