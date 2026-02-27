package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/joseph0x45/surge/internal/models"
)

func (h *Handler) authRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			log.Println("Failed to get cookie:", err.Error())
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		user, err := h.conn.GetUser("id", cookie.Value)
		if err != nil {
			log.Println(err.Error())
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		if user == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(
			r.Context(),
			"user",
			user,
		)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *Handler) RenderLoginPage(w http.ResponseWriter, r *http.Request) {
	if err := h.templates.ExecuteTemplate(w, "login", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	h.render(w, "login", models.PageData{
		Title: "Authenticate",
	})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
