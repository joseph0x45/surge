package handlers

import (
	"net/http"

	"github.com/joseph0x45/surge/internal/models"
)

func (h *Handler) RenderStatsPage(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		http.Redirect(w, r, "/auth", http.StatusSeeOther)
		return
	}
	logs, err := h.conn.GetUserLogs(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.render(w, "stats", map[string]any{
		"Logs": logs,
	})
}
