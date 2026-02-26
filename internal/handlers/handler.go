package handlers

import (
	"html/template"

	"github.com/joseph0x45/surge/internal/db"
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
