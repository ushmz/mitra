package handler

import "context"

type TaskHandler struct{}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{}
}

func (h *TaskHandler) GetTaskByID(ctx context.Context, taskID int) {

}
