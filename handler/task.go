package handler

import (
	"mitra/domain"
	"mitra/store"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
)

type TaskHandler struct {
	Store *store.Store
}

func NewTaskHandler(store *store.Store) *TaskHandler {
	return &TaskHandler{Store: store}
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

type AssignTaskRequest struct {
	domain.RequestBody
	UserID int                  `json:"user_id"`
	Used   domain.TaskTopicUsed `json:"used"`
}

func (h *TaskHandler) AssignTask(w http.ResponseWriter, r *http.Request) {
	if h == nil {
		render.Render(w, r, NewErrResponseRenderer(ErrNilReceiver, http.StatusInternalServerError))
		return
	}

	p := &AssignTaskRequest{}
	if err := render.Bind(r, p); err != nil {
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusBadRequest))
		return
	}

	assigned, err := h.Store.Task.AssignTask(r.Context(), p.UserID, &p.Used)
	if err != nil {
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusInternalServerError))
		return
	}

	render.Render(w, r, NewResponseRenderer(assigned, http.StatusOK))
}

type TaskRequest struct {
	domain.RequestBody
	TaskID int `json:"task_id"`
}

func (h *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	if h == nil {
		render.Render(w, r, NewErrResponseRenderer(ErrNilReceiver, http.StatusInternalServerError))
		return
	}

	ctx := r.Context()

	param := r.URL.Query().Get("tid")
	taskID, err := strconv.ParseInt(param, 0, 64)
	if err != nil {
		render.Render(w, r, NewErrResponseRenderer(ErrInvalidParameter, http.StatusBadRequest))
		return
	}

	task, err := h.Store.Task.GetTask(ctx, taskID)
	if err != nil {
		render.Render(w, r, NewResponseRenderer(ErrInternal, http.StatusInternalServerError))
		return
	}

	render.Render(w, r, NewResponseRenderer(task, http.StatusOK))
}
