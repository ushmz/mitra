package handler

import (
	"fmt"
	"mitra/domain"
	"mitra/store"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
)

type SearchHandler struct {
	Store *store.Store
}

func NewSearchHandler(store *store.Store) *SearchHandler {
	return &SearchHandler{
		Store: store,
	}
}

type ListSearchResultRequest struct {
	domain.RequestBody
	Offset int
	Limit  int
	TaskID int
}

func (h *SearchHandler) GetSimilarweb(w http.ResponseWriter, r *http.Request) {
	p := &ListSearchResultRequest{}

	taskID := r.URL.Query().Get("task")
	if taskID == "" {
		render.Render(w, r, NewErrResponseRenderer(ErrNoRequiredParameter, http.StatusBadRequest))
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
			render.Render(w, r, NewErrResponseRenderer(ErrInvalidParameter, http.StatusBadRequest))
		} else {
			p.Offset = o
		}
	}

	if limit := r.URL.Query().Get("limit"); limit == "" {
		p.Limit = 10
	} else {
		if l, err := strconv.Atoi(limit); err != nil {
			render.Render(w, r, NewErrResponseRenderer(ErrInvalidParameter, http.StatusBadRequest))
		} else {
			p.Limit = l
		}
	}

	ctx := r.Context()
	taskIDs := []int{52, 53, 54, 55}
	rs, err := h.Store.Search.GetSimilarwebPagesByPageIDs(ctx, taskIDs)
	if err != nil {
		fmt.Println(err)
		render.Render(w, r, NewErrResponseRenderer(ErrInternal, http.StatusInternalServerError))
		return
	}

	render.Render(w, r, NewResponseRenderer(rs, http.StatusOK))
}

// ListSearchResult return listed search result
func (h *SearchHandler) ListSearchResult(w http.ResponseWriter, r *http.Request) {
	p := &ListSearchResultRequest{}

	taskID := r.URL.Query().Get("task")
	if taskID == "" {
		render.Render(w, r, NewErrResponseRenderer(ErrNoRequiredParameter, http.StatusBadRequest))
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
			render.Render(w, r, NewErrResponseRenderer(ErrInvalidParameter, http.StatusBadRequest))
		} else {
			p.Offset = o
		}
	}

	if limit := r.URL.Query().Get("limit"); limit == "" {
		p.Limit = 10
	} else {
		if l, err := strconv.Atoi(limit); err != nil {
			render.Render(w, r, NewErrResponseRenderer(ErrInvalidParameter, http.StatusBadRequest))
		} else {
			p.Limit = l
		}
	}

	result, err := h.Store.Search.GetSearchPages(r.Context(), p.TaskID, p.Offset, p.Limit, 10)
	if err != nil {
		fmt.Println(err)
		render.Render(w, r, NewErrResponseRenderer(ErrInternal, http.StatusInternalServerError))
		return
	}

	render.Render(w, r, NewResponseRenderer(result, http.StatusOK))
}
