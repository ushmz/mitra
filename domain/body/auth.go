package body

import "net/http"

type RegisterRequest struct {
	UID string `json:"uid"`
}

func (p *RegisterRequest) Bind(r *http.Request) error {
	if p == nil {
		return ErrBadRequest
	}
	return nil
}
