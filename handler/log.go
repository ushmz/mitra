package handler

import (
	"fmt"
	"mitra/domain"
	"mitra/store"
	"net/http"

	"github.com/go-chi/render"
)

type LogHandler struct {
	Store *store.Store
}

func NewLogHandler(store *store.Store) *LogHandler {
	return &LogHandler{Store: store}
}

// DwellTimeLogRequest : Struct for SERP dwell time log request body
type DwellTimeLogRequest struct {
	domain.RequestBody
	// UserID : The ID of user (worker)
	UserID int `json:"user"`

	// TaskID : The ID of task that user working.
	TaskID int `json:"task"`

	// Condition : User's condition ID that means group and task category.
	Condition string `json:"condition"`
}

// CreateDwellTimeLog creates new dwell time log entity.
// It is assumed that this will be called once a second.
func (h *LogHandler) CreateDwellTimeLog(w http.ResponseWriter, r *http.Request) {
	if h == nil {
		render.Render(w, r, NewErrResponseRenderer(ErrNilReceiver, http.StatusInternalServerError))
	}

	ctx := r.Context()

	p := &DwellTimeLogRequest{}
	if err := render.Bind(r, p); err != nil {
		fmt.Println(err)
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusBadRequest))
		return
	}

	m := domain.DwellTimeLog{
		UserID:    p.UserID,
		TaskID:    p.TaskID,
		Condition: p.Condition,
	}

	if err := h.Store.Log.CreateDwelltimeLog(ctx, m); err != nil {
		fmt.Println(err)
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusInternalServerError))
		return
	}

	render.Render(w, r, NewResponseRenderer(nil, http.StatusNoContent))
}

// ClickLogRequest : Struct for SERP click log request body
type ClickLogRequest struct {
	domain.RequestBody
	// Uid : The ID of user (worker)
	UID int `json:"user"`

	// TaskID : The ID of task that user working.
	TaskID int `json:"task"`

	// ConditionID : User's condition ID that means group and task category.
	Condition string `json:"condition"`

	// Time : User's page viewing time.
	Time int `json:"time"`

	// Rank : Search result rank that user clicked.
	Rank int `json:"rank"`

	// IsVisible : Risk is visible or not.
	IsVisible bool `json:"visible"`

	// IsFirstClick : The click event is the first click or not
	IsFirstClick bool `json:"is_first"`
}

// CreateClickLog creates new click log entity.
func (h *LogHandler) CreateClickLog(w http.ResponseWriter, r *http.Request) {
	if h == nil {
		render.Render(w, r, NewErrResponseRenderer(ErrNilReceiver, http.StatusInternalServerError))
	}

	ctx := r.Context()

	p := &ClickLogRequest{}
	if err := render.Bind(r, p); err != nil {
		fmt.Println(err)
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusBadRequest))
		return
	}

	l := domain.ClickLog{
		UserID:       p.UID,
		TaskID:       p.TaskID,
		Condition:    p.Condition,
		Time:         p.Time,
		Rank:         p.Rank,
		IsVisible:    p.IsVisible,
		IsFirstClick: p.IsFirstClick,
		Event:        "click",
	}

	if err := h.Store.Log.CreateClickLog(ctx, l); err != nil {
		fmt.Println(err)
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusInternalServerError))
		return
	}

	render.Render(w, r, NewResponseRenderer(nil, http.StatusNoContent))
}

// CreateHoverLog creates new hover log entity.
// func (c *LogHandler) CreateHoverLog(w http.ResponseWriter, r *http.Request) {
// 	p := &ClickLogRequest{}
// 	if err := render.Bind(r, p); err != nil {
// 		render.Render(w, r, NewErrResponseRenderer(err, http.StatusBadRequest))
// 		return
// 	}
//
// 	render.Render(w, r, NewResponseRenderer(p, http.StatusOK))
// }
