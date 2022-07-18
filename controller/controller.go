package controller

import (
	"mitra/config"
	"net/http"

	"github.com/go-chi/render"
)

// ErrResponse is the struct used with error response
type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	Message string `json:"message"`
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

// Index return greeting response
func Index(w http.ResponseWriter, r *http.Request) {
	c := config.GetConfig()
	v := c.GetString("version")
	if v == "" {
		v = "beta"
	}
	greet := "" +
		`   __  ____ __           ` + "\n" +
		`  /  |/  (_) /________ _ ` + "\n" +
		` / /|_/ / / __/ __/ _ \/ ` + "\n" +
		`/_/  /_/_/\__/_/  \_,_/  ` + c.GetString("version") + "\n" +
		`Simple backend API server for the user-study.` + "\n" +
		`=================================================` + "\n"
	render.PlainText(w, r, greet)
}
