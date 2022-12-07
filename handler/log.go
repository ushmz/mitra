package handler

import (
	"fmt"
	"mitra/domain"
	"net/http"

	"github.com/go-chi/render"
)

type LogHandler struct{}

func NewLogHandler() *LogHandler {
	return &LogHandler{}
}

// DwellTimeLogRequest : Struct for SERP dwell time log request body
type DwellTimeLogRequest struct {
	domain.RequestBody
	// UserID : The ID of user (worker)
	UserID int `db:"user_id" json:"user"`

	// TaskID : The ID of task that user working.
	TaskID int `db:"task_id" json:"task"`

	// ConditionID : User's condition ID that means group and task category.
	ConditionID int `db:"condition_id" json:"condition"`
}

// CreateDwellTimeLog creates new dwell time log entity.
// It is assumed that this will be called once a second.
func (c *LogHandler) CreateDwellTimeLog(w http.ResponseWriter, r *http.Request) {
	p := &DwellTimeLogRequest{}
	if err := render.Bind(r, p); err != nil {
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusBadRequest))
		return
	}

	m := domain.DwellTimeLog{
		UserID:      p.UserID,
		TaskID:      p.TaskID,
		ConditionID: p.ConditionID,
	}
	fmt.Println("Store log", m)
	render.Render(w, r, NewResponseRenderer(p, http.StatusOK))
}

// ClickLogRequest : Struct for SERP click log request body
type ClickLogRequest struct {
	domain.RequestBody
	// Uid : The ID of user (worker)
	UID int `db:"user_id" json:"user"`

	// TaskID : The ID of task that user working.
	TaskID int `db:"task_id" json:"task"`

	// ConditionID : User's condition ID that means group and task category.
	ConditionID int `db:"condition_id" json:"condition"`

	// Time : User's page viewing time.
	Time int `db:"time_on_page" json:"time"`

	// Page : The ID of page that user clicked.
	PageID int `db:"page_id" json:"page_id"`

	// Rank : Search result rank that user clicked.
	Rank int `db:"page_rank" json:"rank"`

	// IsVisible : Risk is visible or not.
	IsVisible bool `db:"is_visible" json:"visible"`

	// IsFirstClick : The click event is the first click or not
	IsFirstClick bool `db:"is_first_click" json:"is_first"`
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
