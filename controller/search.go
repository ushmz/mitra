package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
)

// ListSearchResult return listed search result
func ListSearchResult(w http.ResponseWriter, r *http.Request) {
	p := &ListSearchResultRequest{}

	if taskID := r.URL.Query().Get("task"); taskID == "" {
		render.Render(w, r, NewErrResponseRenderer(errors.New("No Required parameter"), http.StatusBadRequest))
	} else {
		if t, err := strconv.Atoi(taskID); err != nil {
			p.TaskID = t
		} else {
			render.Render(w, r, NewErrResponseRenderer(err, http.StatusBadRequest))
		}
	}

	if offset := r.URL.Query().Get("offset"); offset == "" {
		p.Offset = 0
	} else {
		if o, err := strconv.Atoi(offset); err != nil {
			p.Offset = o
		} else {
			render.Render(w, r, NewErrResponseRenderer(err, http.StatusBadRequest))
		}
	}

	if attribute := r.URL.Query().Get("attr"); attribute == "1" {
		p.Attribute = true
	} else {
		p.Attribute = false
	}

	render.Render(w, r, NewResponseRenderer(p, http.StatusOK))
}
