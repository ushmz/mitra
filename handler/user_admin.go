package handler

import (
	"mitra/domain"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type AdminUserHandler struct{}

func NewAdminUserHandler() *AdminUserHandler {
	return &AdminUserHandler{}
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
func (h *AdminUserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	data := []domain.User{
		{UID: "mitra"},
		{UID: "varuna"},
	}
	render.JSON(w, r, data)
}

// GetUser get specific user information
func (h *AdminUserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	render.JSON(w, r, userID)
}

// DeleteUser deletes user information
func (h *AdminUserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	render.JSON(w, r, userID)
}
