package handler

import (
	"mitra/domain/model"
	"net/http"

	"github.com/go-chi/render"
)

type LogHandler struct{}

func NewLogHandler() *LogHandler {
	return &LogHandler{}
}

// CreateDwellTimeLog creates new dwell time log entity.
// It is assumed that this will be called once a second.
func (c *LogHandler) CreateDwellTimeLog(w http.ResponseWriter, r *http.Request) {
	p := &DwellTimeLogRequest{}
	if err := render.Bind(r, p); err != nil {
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusBadRequest))
		return
	}

	m := model.DwellTimeLog{
		UserID:      p.UserID,
		TaskID:      p.TaskID,
		ConditionID: p.ConditionID,
	}
	render.Render(w, r, NewResponseRenderer(p, http.StatusOK))
}

// CreateClickLog creates new click log entity.
func (c *LogHandler) CreateClickLog(w http.ResponseWriter, r *http.Request) {
	p := &ClickLogRequest{}
	if err := render.Bind(r, p); err != nil {
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusBadRequest))
		return
	}

	render.Render(w, r, NewResponseRenderer(p, http.StatusOK))
}

// CreateHoverLog creates new hover log entity.
func (c *LogHandler) CreateHoverLog(w http.ResponseWriter, r *http.Request) {
	p := &ClickLogRequest{}
	if err := render.Bind(r, p); err != nil {
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusBadRequest))
		return
	}

	render.Render(w, r, NewResponseRenderer(p, http.StatusOK))
}
