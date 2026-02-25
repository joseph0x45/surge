package handler

import (
	"daemon/db"
	"html/template"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	templates *template.Template
	conn      *db.Conn
	version   string
}

func NewHandler(
	templates *template.Template,
	conn *db.Conn,
	version string,
) *Handler {
	return &Handler{
		templates: templates,
		conn:      conn,
		version:   version,
	}
}

func (h *Handler) RegisterRoutes(r chi.Router) {

}
