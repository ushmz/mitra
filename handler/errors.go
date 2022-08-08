package handler

import "errors"

var (
	// ErrNilReceiver means the method is called with Nil receiver
	ErrNilReceiver = errors.New("Called with Nil receiver")

	// ErrBadRequest means the function is called with no required parameters
	ErrBadRequest = errors.New("No required parameter")

	// ErrNotFound means the requested resource is not found
	ErrNotFound = errors.New("Resource not found")

	// ErrInternal means the errors which details does not matter
	ErrInternal = errors.New("Somethig went wrong")
)
