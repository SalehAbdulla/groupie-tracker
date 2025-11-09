package handlers

import (
	"bytes"
	"groupie-tracker/internal/constants"
	"groupie-tracker/ui"
	"html/template"
	"io/fs"
	"net/http"
)

type Handlers struct {
	base    *template.Template 
	Static  fs.FS
	Artists []constants.ArtistView
}

func New(view []constants.ArtistView) (*Handlers, error) {
	t, err := template.ParseFS(ui.Files, "templates/*.html")
	if err != nil {
		return nil, constants.InternalServerError
	}
	sub, _ := fs.Sub(ui.Files, "templates")
	return &Handlers{base: t, Static: sub, Artists: view}, nil
}

func (h *Handlers) cloneBase() (*template.Template, error) {
	return h.base.Clone() 
}

func (h *Handlers) render(w http.ResponseWriter, name string, data any) {
	t, err := h.cloneBase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = buf.WriteTo(w)
}
