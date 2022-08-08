package handler

import (
	"mitra/domain/model"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type AdminHandler struct{}

func NewAdminHandler() *AdminHandler {
	return &AdminHandler{}
}

type SignInRequest struct {
	Email  string `json:"email"`
	Passwd string `json:"password"`
}

func (p *SignInRequest) Bind(r *http.Request) error {
	if p == nil {
		return ErrBadRequest
	}
	return nil
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	p := &SignInRequest{}
	if err := render.Bind(r, p); err != nil {
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusBadRequest))
	}
	render.JSON(w, r, p.Email)
}

// ListUsers lists all attendees
func (h *AdminHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	data := []model.User{
		{UID: "mitra"},
		{UID: "varuna"},
	}
	render.JSON(w, r, data)
}

// GetUser get specific user information
func (h *AdminHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	render.JSON(w, r, userID)
}

// DeleteUser deletes user information
func (h *AdminHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	render.JSON(w, r, userID)
}

// ListDwellTimeLogs lists all dwell time logs
func (h *AdminHandler) ListDwellTimeLogs(w http.ResponseWriter, r *http.Request) {
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
func (h *AdminHandler) ListClickLogs(w http.ResponseWriter, r *http.Request) {
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
func (h *AdminHandler) ListHoverLogs(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")

	if userID == "" {
		// Get all logs
		render.JSON(w, r, "UserID is empty: "+userID)
		return
	}

	// Get specific user's log
	render.JSON(w, r, userID)
}
