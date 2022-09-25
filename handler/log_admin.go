package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type AdminLogHandler struct{}

func NewAdminLogHandler() *AdminLogHandler {
	return &AdminLogHandler{}
}

// ListDwellTimeLogs lists all dwell time logs
func (h *AdminLogHandler) ListDwellTimeLogs(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")

	if userID == "" {
		// Get all logs
		render.JSON(w, r, "UserID is empty: "+userID)
		return
	}

	// Get specific user's log
	render.JSON(w, r, userID)
}

// ListClickLogs lists all click logs
func (h *AdminLogHandler) ListClickLogs(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")

	if userID == "" {
		// Get all logs
		render.JSON(w, r, "UserID is empty: "+userID)
		return
	}

	// Get specific user's log
	render.JSON(w, r, userID)
}

// ListHoverLogs lists all hover logs
func (h *AdminLogHandler) ListHoverLogs(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")

	if userID == "" {
		// Get all logs
		render.JSON(w, r, "UserID is empty: "+userID)
		return
	}

	// Get specific user's log
	render.JSON(w, r, userID)
}
