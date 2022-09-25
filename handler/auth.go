package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

type RegisterRequest struct {
	UID string `json:"uid"`
}

func (p *RegisterRequest) Bind(r *http.Request) error {
	if p == nil {
		return ErrBadRequest
	}
	return nil
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	p := &RegisterRequest{}
	if err := render.Bind(r, p); err != nil {
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusBadRequest))
	}
	render.Render(w, r, NewResponseRenderer("", http.StatusOK))
}
