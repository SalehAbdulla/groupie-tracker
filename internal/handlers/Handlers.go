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
		return nil, constants.ErrInternalServer
	}
	sub, _ := fs.Sub(ui.Files, "templates")
	return &Handlers{base: t, Static: sub, Artists: view}, nil
}

func (h *Handlers) cloneBase() (*template.Template, error) {
	return h.base.Clone()
}

func (h *Handlers) Render(w http.ResponseWriter, r *http.Request, name string, data any) {
	t, err := h.cloneBase()
	if err != nil {
		h.InternalServerError(w, r)
		return
	}

	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, name, data); err != nil {
		h.InternalServerError(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = buf.WriteTo(w)
}
