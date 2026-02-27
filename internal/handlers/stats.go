package handlers

import (
	"net/http"

	"github.com/joseph0x45/surge/internal/models"
)

func (h *Handler) RenderStatsPage(w http.ResponseWriter, r *http.Request) {
	h.render(w, "stats", models.PageData{
		Title: "Stats",
	})
}
