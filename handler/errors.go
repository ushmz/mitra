package handler

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

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

// ErrResponse is the struct used with error response
type ErrResponse struct {
	Err            error  `json:"-"`
	HTTPStatusCode int    `json:"-"`
	Message        string `json:"message"`
}

// Render renders response header and body
func (er ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, er.HTTPStatusCode)
	return nil
}

// NewErrResponseRenderer return the struct for the accerate response
func NewErrResponseRenderer(err error, statusCode int) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: statusCode,
		Message:        err.Error(),
	}
}

// AccurateResponse is the struct used with error response
type AccurateResponse struct {
	HTTPStatusCode int `json:"-"`

	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

// Render renders response header and body
func (ar AccurateResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, ar.HTTPStatusCode)
	return nil
}

// NewResponseRenderer return the struct for the accerate response
func NewResponseRenderer(data interface{}, statusCode int) render.Renderer {
	msg := ""
	if data == nil {
		msg = http.StatusText(statusCode)
	}
	return &AccurateResponse{
		HTTPStatusCode: statusCode,
		Data:           data,
		Message:        msg,
	}
}
