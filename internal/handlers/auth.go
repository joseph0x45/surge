package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/joseph0x45/goutils"
	"github.com/joseph0x45/surge/internal/models"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func (h *Handler) authRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			log.Println("Failed to get cookie:", err.Error())
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			return
		}
		session, err := h.conn.GetSession(cookie.Value)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			return
		}
		if session == nil {
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			return
		}
		user, err := h.conn.GetUser("id", session.UserID)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			return
		}
		if user == nil {
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
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
	if err := h.templates.ExecuteTemplate(w, "auth", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	user, err := h.conn.GetUser("username", username)
	if err != nil {
		log.Println(err)
		h.render(w, "auth", map[string]string{
			"Error": "Something went wrong",
		})
		return
	}
	if user == nil {
		h.render(w, "auth", map[string]string{
			"Error": fmt.Sprintf("User %s not found\n", username),
		})
		return
	}
	if !goutils.HashMatchesPassword(user.Password, password) {
		h.render(w, "auth", map[string]string{
			"Error": "Invalid credentials",
		})
		return
	}
	session := models.Session{
		ID:     gonanoid.Must(),
		UserID: user.ID,
	}
	if err := h.conn.InsertSession(&session); err != nil {
		log.Println(err)
		h.render(w, "auth", map[string]string{
			"Error": "Something went wrong",
		})
		return
	}
	cookie := &http.Cookie{
		Name:     "session",
		Value:    session.ID,
		Path:     "/",
		HttpOnly: true,
		Secure:   h.version != "debug",
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(100 * 365 * 24 * time.Hour),
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
