package body

import (
	"net/http"
)

type CreateUserRequest struct {
	UID string `json:"uid"`
}

func (p *CreateUserRequest) Bind(r *http.Request) error {
	if p == nil {
		return ErrBadRequest
	}
	return nil
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
