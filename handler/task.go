package handler

import (
	"mitra/domain"
	"mitra/store"
	"net/http"

	"github.com/go-chi/render"
)

type TaskHandler struct {
	Store *store.Store
}

func NewTaskHandler(store *store.Store) *TaskHandler {
	return &TaskHandler{Store: store}
}

type AssignTaskRequest struct {
	domain.RequestBody
	UserID int `json:"user_id"`
}

func (h *TaskHandler) GetTaskQueries(w http.ResponseWriter, r *http.Request) {
	if h == nil {
		render.Render(w, r, NewErrResponseRenderer(nil, http.StatusInternalServerError))
		return
	}

	queries, err := h.Store.Task.GetTaskQueries(r.Context())
	if err != nil {
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusInternalServerError))
		return
	}

	render.Render(w, r, NewResponseRenderer(queries, http.StatusOK))
}

func (h *TaskHandler) AssignTask(w http.ResponseWriter, r *http.Request) {
	if h == nil {
		render.Render(w, r, NewErrResponseRenderer(nil, http.StatusInternalServerError))
		return
	}

	p := &AssignTaskRequest{}
	if err := render.Bind(r, p); err != nil {
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusBadRequest))
		return
	}

	assigned, err := h.Store.Task.AssignTask(r.Context(), p.UserID, []string{})
	if err != nil {
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusInternalServerError))
		return
	}

	render.Render(w, r, NewResponseRenderer(assigned, http.StatusOK))
}

// func (h *TaskHandler) GetTaskByID(ctx context.Context, taskID int) error { }
