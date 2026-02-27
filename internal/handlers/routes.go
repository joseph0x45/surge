package handlers

import "github.com/go-chi/chi/v5"

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Use(h.authRequired)
		r.Get("/", h.RenderApp)
		r.Post("/sync", h.Sync)
		r.Get("/stats", h.RenderStatsPage)
	})

	r.Get("/auth", h.RenderLoginPage)

	r.Post("/auth", h.Auth)
}
