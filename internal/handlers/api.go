package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/joseph0x45/surge/internal/models"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func (h *Handler) Sync(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*models.User)
	if !ok {
		http.Redirect(w, r, "/auth", http.StatusSeeOther)
		return
	}
	payload := &struct {
		Elapsed int `json:"elapsed"`
	}{}
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		log.Println("Error while decoding data", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	newLog := models.Log{
		ID:        gonanoid.Must(),
		UserID:    user.ID,
		DateStr:   time.Now().Format("Monday, Jan 2"),
		Elapsed:   payload.Elapsed,
		CreatedAt: time.Now().Unix(),
	}
	if err := h.conn.UpdateLogs(&newLog); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
