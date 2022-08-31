package store

import (
	"context"
	"mitra/domain/model"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

// TaskStore : The store object for task data source.
type TaskStore interface {
	CreateTask(ctx context.Context, name string, description string) (*model.Task, error)
	GetTaskSimple(ctx context.Context, id int64) (*model.TaskSimple, error)
	GetTask(ctx context.Context, id int64) (*model.Task, error)
	ListTasks(ctx context.Context, limit, offset int) ([]model.Task, error)
	UpdateTasks(ctx context.Context, task *model.Task) (*model.Task, error)
	DeleteTasks(ctx context.Context, id int64) error
}

// TaskStoreImpl : Implementation of TaskStore interface.
type TaskStoreImpl struct {
	db *sqlx.DB
}

// NewTaskStore returns new TaskStore implementation.
func NewTaskStore(db *sqlx.DB) TaskStore {
	return &TaskStoreImpl{db: db}
}

// CreateTask creates a new task record in DB
func (s *TaskStoreImpl) CreateTask(ctx context.Context, title, description string) (*model.Task, error) {
	if s == nil {
		return nil, ErrNilReceiver
	}

	q, a, err := dialect.
		Insert("tasks").
		Rows(goqu.Record{"title": title, "description": description}).
		ToSQL()
	if err != nil {
		return nil, ErrQueryBuildFailure
	}

	rs, err := s.db.ExecContext(ctx, q, a...)
	if err != nil {
		return nil, ErrDatabaseExecutionFailere
	}

	id, err := rs.LastInsertId()
	if err != nil {
		return nil, ErrDatabaseExecutionFailere
	}

	return &model.Task{
		ID:          id,
		Title:       title,
		Description: description,
	}, nil
}

// GetTaskSimple gets single instance of task for normal users.
func (s *TaskStoreImpl) GetTaskSimple(ctx context.Context, id int64) (*model.TaskSimple, error) {
	if s == nil {
		return nil, ErrNilReceiver
	}

	q, a, err := dialect.Select("title", "description").From("tasks").ToSQL()
	if err != nil {
		return nil, ErrQueryBuildFailure
	}

	rs := &model.TaskSimple{}
	if err := s.db.GetContext(ctx, rs, q, a...); err != nil {
		return nil, ErrDatabaseExecutionFailere
	}

	return rs, nil
}

// GetTask gets single instance of task for admin users.
func (s *TaskStoreImpl) GetTask(ctx context.Context, id int64) (*model.Task, error) {
	if s == nil {
		return nil, ErrNilReceiver
	}

	q, a, err := dialect.Select("*").From("tasks").ToSQL()
	if err != nil {
		return nil, ErrQueryBuildFailure
	}

	rs := &model.Task{}
	if err := s.db.GetContext(ctx, rs, q, a...); err != nil {
		return nil, ErrDatabaseExecutionFailere
	}

	return rs, nil
}

// ListTasks gets a number of tasks for admin users.
func (s *TaskStoreImpl) ListTasks(ctx context.Context, limit int, offset int) ([]model.Task, error) {
	if s == nil {
		return nil, ErrNilReceiver
	}

	b := dialect.Select("*").From("tasks")

	if limit > 0 {
		b = b.Limit(uint(limit))
	}

	if offset > 0 {
		b = b.Offset(uint(offset))
	}

	q, a, err := b.ToSQL()
	if err != nil {
		return nil, ErrQueryBuildFailure
	}

	rs := make([]model.Task, limit)
	if err := s.db.SelectContext(ctx, rs, q, a...); err != nil {
		return nil, ErrDatabaseExecutionFailere
	}

	return rs, nil
}

// UpdateTasks updates a task for admin users.
func (s *TaskStoreImpl) UpdateTasks(ctx context.Context, task *model.Task) (*model.Task, error) {
	if s == nil {
		return nil, ErrNilReceiver
	}

	q, a, err := dialect.Update("tasks").Set(task).ToSQL()
	if err != nil {
		return nil, ErrQueryBuildFailure
	}

	if _, err := s.db.ExecContext(ctx, q, a...); err != nil {
		return nil, ErrDatabaseExecutionFailere
	}

	return task, nil

}

// DeleteTasks deletes a task for admin users.
func (s *TaskStoreImpl) DeleteTasks(ctx context.Context, id int64) error {
	if s == nil {
		return ErrNilReceiver
	}

	q, a, err := dialect.Delete("tasks").Where(goqu.Ex{"id": id}).ToSQL()
	if err != nil {
		return ErrQueryBuildFailure
	}

	if _, err := s.db.ExecContext(ctx, q, a...); err != nil {
		return ErrDatabaseExecutionFailere
	}

	return nil
}
