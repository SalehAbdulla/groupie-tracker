package handlers

import (
	"html/template"
	"net/http"
)

type Handlers struct {
	templatesPath string
}

func New(templatesPath string) *Handlers {
	return &Handlers{templatesPath: templatesPath}
}

func (h *Handlers) render(w http.ResponseWriter, name string, data any) {
	t, err := template.ParseFiles(h.templatesPath + "/" + name)
	if err != nil {http.Error(w, err.Error(), http.StatusInternalServerError); return}
	_ = t.Execute(w, data)
}
