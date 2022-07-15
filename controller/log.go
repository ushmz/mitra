package controller

import (
	"net/http"

	"github.com/go-chi/render"
)

// CreateDwellTimeLog creates new dwell time log entity.
// It is assumed that this will be called once a second.
func CreateDwellTimeLog(w http.ResponseWriter, r *http.Request) {
	p := &DwellTimeLogRequest{}
	if err := render.Bind(r, p); err != nil {
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusBadRequest))
		return
	}
	render.Render(w, r, NewResponseRenderer(p, http.StatusOK))
}

// CreateClickLog creates new click log entity.
func CreateClickLog(w http.ResponseWriter, r *http.Request) {
	p := &ClickLogRequest{}
	if err := render.Bind(r, p); err != nil {
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusBadRequest))
		return
	}

	render.Render(w, r, NewResponseRenderer(p, http.StatusOK))
}

// CreateHoverLog creates new click log entity.
func CreateHoverLog(w http.ResponseWriter, r *http.Request) {
	p := &ClickLogRequest{}
	if err := render.Bind(r, p); err != nil {
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusBadRequest))
		return
	}

	render.Render(w, r, NewResponseRenderer(p, http.StatusOK))
}
