package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
)

type SearchHandler struct{}

func NewSearchHandler() *SearchHandler {
	return &SearchHandler{}
}

// ListSearchResult return listed search result
func (h *SearchHandler) ListSearchResult(w http.ResponseWriter, r *http.Request) {
	p := &ListSearchResultRequest{}

	taskID := r.URL.Query().Get("task")
	if taskID == "" {
		render.Render(w, r, NewErrResponseRenderer(
			errors.New("No Required parameter"),
			http.StatusBadRequest,
		))
		return
	}

	if t, err := strconv.Atoi(taskID); err != nil {
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusBadRequest))
	} else {
		p.TaskID = t
	}

	if offset := r.URL.Query().Get("offset"); offset == "" {
		p.Offset = 0
	} else {
		if o, err := strconv.Atoi(offset); err != nil {
			render.Render(w, r, NewErrResponseRenderer(err, http.StatusBadRequest))
		} else {
			p.Offset = o
		}
	}

	if attribute := r.URL.Query().Get("attr"); attribute == "1" {
		p.Attribute = true
	} else {
		p.Attribute = false
	}

	render.Render(w, r, NewResponseRenderer(p, http.StatusOK))
}
