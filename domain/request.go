package domain

import "net/http"

type RequestBody struct{}

// Bind binds request body to the receiver
func (b *RequestBody) Bind(r *http.Request) error {
	if b == nil {
		return ErrInvalidRequestBody
	}
	return nil
}
