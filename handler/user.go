package handler

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/render"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

type CreateUserRequest struct {
	UID string `json:"uid"`
}

func (p *CreateUserRequest) Bind(r *http.Request) error {
	if p == nil {
		return ErrBadRequest
	}
	return nil
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	p := &CreateUserRequest{}
	if err := render.Bind(r, p); err != nil {
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusBadRequest))
	}
	render.Render(w, r, NewResponseRenderer("", http.StatusOK))
}

type IssueCompletionCodeRequest struct {
	UserID int `db:"user_id"`
}

func (p *IssueCompletionCodeRequest) Bind(r *http.Request) error {
	if p == nil {
		return ErrBadRequest
	}
	return nil
}

func (h *UserHandler) IssueCompletionCode(w http.ResponseWriter, r *http.Request) {
	req := &IssueCompletionCodeRequest{}
	if err := render.Bind(r, req); err != nil {
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusBadRequest))
		return
	}
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(100000)

	// p := model.CompletionCode{UserID: req.UserID, Code: code}

	render.Render(w, r, NewResponseRenderer(code, http.StatusOK))
}
