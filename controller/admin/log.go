package admin

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// ListDwellTimeLogs lists all dwell time logs
func ListDwellTimeLogs(w http.ResponseWriter, r *http.Request) {
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
func ListClickLogs(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")

	if userID == "" {
		// Get all logs
		render.JSON(w, r, "UserID is empty: "+userID)
		return
	}

	// Get specific user's log
	render.JSON(w, r, userID)
}
