package body

import (
	"net/http"
)

type ListSearchResultRequest struct {
	Offset    int  `param:"offset"`
	TaskID    int  `param:"task"`
	Attribute bool `param:"attr"`
}

func (p *ListSearchResultRequest) Bind(r *http.Request) error {
	if p == nil {
		return ErrBadRequest
	}
	return nil
}
