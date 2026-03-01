package handlers

import (
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

func (h *Handler) render(w http.ResponseWriter, templateName string, data any) {
	if err := h.templates.ExecuteTemplate(w, templateName, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) RenderApp(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		http.Redirect(w, r, "/auth", http.StatusSeeOther)
		return
	}
	if err := h.templates.ExecuteTemplate(w, "main", map[string]any{
		"User": user,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
