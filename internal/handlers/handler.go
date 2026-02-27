package handlers

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/joseph0x45/surge/internal/db"
	"github.com/joseph0x45/surge/internal/models"
)

type Handler struct {
	conn      *db.Conn
	templates *template.Template
	version   string
}

func NewHandler(
	conn *db.Conn,
	templates *template.Template,
	version string,
) *Handler {
	return &Handler{
		conn:      conn,
		templates: templates,
		version:   version,
	}
}

func (h *Handler) render(
	w http.ResponseWriter, templateName string,
	data models.PageData,
) {
	var contentBuffer bytes.Buffer
	if err := h.templates.ExecuteTemplate(
		&contentBuffer, templateName, data,
	); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data.Content = template.HTML(contentBuffer.String())
	var responseBuffer bytes.Buffer
	if err := h.templates.ExecuteTemplate(
		&responseBuffer, "base", data,
	); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseBuffer.Bytes())
}

func (h *Handler) RenderApp(w http.ResponseWriter, r *http.Request) {
	h.render(w, "index", models.PageData{
		Title: "Surge",
	})
}
