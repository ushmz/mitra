package domain

import "errors"

var (
	ErrInvalidRequestBody = errors.New("Cannot read request body")
)
