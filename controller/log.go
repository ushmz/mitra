package controller

import (
	"net/http"

	"github.com/go-chi/render"
)

func CreateDwellTimeLog(w http.ResponseWriter, r *http.Request) {
	p := &DwellTimeLogRequest{}
	if err := render.Bind(r, p); err != nil {
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusBadRequest))
		return
	}
	render.Render(w, r, NewResponseRenderer(p, http.StatusOK))
}

func CreateClickLog(w http.ResponseWriter, r *http.Request) {
	p := &ClickLogRequest{}
	if err := render.Bind(r, p); err != nil {
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusBadRequest))
		return
	}

	render.Render(w, r, NewResponseRenderer(p, http.StatusOK))
}
func CreateHoverLog(w http.ResponseWriter, r *http.Request) {
	p := &ClickLogRequest{}
	if err := render.Bind(r, p); err != nil {
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusBadRequest))
		return
	}

	render.Render(w, r, NewResponseRenderer(p, http.StatusOK))
}
