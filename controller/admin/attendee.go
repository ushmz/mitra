package admin

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// Attendee shows attendees' information
type Attendee struct {
	UID string `json:"uid"`
}

// ListAttendees lists all attendees
func ListAttendees(w http.ResponseWriter, r *http.Request) {
	data := []Attendee{
		{UID: "mitra"},
		{UID: "varuna"},
	}
	render.JSON(w, r, data)
}

// GetAttendees get specific user information
func GetAttendees(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	render.JSON(w, r, userID)
}

// DeleteAttendees deletes user information
func DeleteAttendees(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	render.JSON(w, r, userID)
}
