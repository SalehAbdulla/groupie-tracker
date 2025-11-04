package handlers

import (
	"bytes"
	"groupie-tracker/internal/constants"
	"groupie-tracker/ui"
	"html/template"
	"io/fs"
	"net/http"
	"strings"
)

type Handlers struct {
	base    *template.Template // never execute this one
	Static  fs.FS
	Artists []constants.ArtistView
}

func New(view []constants.ArtistView) *Handlers {
	t := template.Must(template.ParseFS(ui.Files, "templates/*.html"))
	sub, _ := fs.Sub(ui.Files, "templates")
	return &Handlers{base: t, Static: sub, Artists: view}
}

func (h *Handlers) cloneBase() (*template.Template, error) {
	return h.base.Clone() // safe because h.base is never executed
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

func (h *Handlers) RenderStr(w http.ResponseWriter, name string, data any) {
	const inlinePrefix = "inline:"
	if strings.HasPrefix(name, inlinePrefix) {
		src := strings.TrimPrefix(name, inlinePrefix)

		t, err := h.cloneBase()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		const inlineName = "__inline__"
		if _, err := t.New(inlineName).Parse(src); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var buf bytes.Buffer
		if err := t.ExecuteTemplate(&buf, inlineName, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = buf.WriteTo(w)
		return
	}

	// fall back to file-backed template by name
	h.render(w, name, data)
}
