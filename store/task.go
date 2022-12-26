package store

import (
	"context"
	"mitra/domain"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/jmoiron/sqlx"
)

// TaskStore : The store object for task data source.
type TaskStore interface {
	GetTaskQueries(ctx context.Context) ([]domain.TaskTopic, error)
	AssignTask(ctx context.Context, userID int, used *domain.TaskTopicUsed) (*domain.AssignedTask, error)
	CreateTask(ctx context.Context, name string, description string) (*domain.Task, error)
	GetTaskSimple(ctx context.Context, id int64) (*domain.TaskSimple, error)
	GetTask(ctx context.Context, id int64) (*domain.Task, error)
	ListTasks(ctx context.Context, limit, offset int) ([]domain.Task, error)
	UpdateTasks(ctx context.Context, task *domain.Task) (*domain.Task, error)
	DeleteTasks(ctx context.Context, id int64) error
	CreateAnswer(ctx context.Context, answer *domain.Answer) error
}

// TaskStoreImpl : Implementation of TaskStore interface.
type TaskStoreImpl struct {
	db *sqlx.DB
}

// NewTaskStore returns new TaskStore implementation.
func NewTaskStore(db *sqlx.DB) TaskStore {
	return &TaskStoreImpl{db: db}
}

func (s *TaskStoreImpl) GetTaskQueries(ctx context.Context) ([]domain.TaskTopic, error) {
	if s == nil {
		return nil, ErrNilReceiver
	}

	q, a, err := dialect.Select("id", "topic").From(goqu.T("tasks")).ToSQL()
	if err != nil {
		return nil, ErrQueryBuildFailure
	}

	dest := []domain.TaskTopic{}
	if err := s.db.SelectContext(ctx, &dest, q, a...); err != nil {
		return nil, ErrDatabaseExecutionFailere
	}

	return dest, nil
}

func (s *TaskStoreImpl) getAssignedTask(ctx context.Context, userID int) (*domain.AssignedTask, error) {
	if s == nil {
		return nil, ErrNilReceiver
	}

	q, a, err := dialect.
		Select("task_id", "condition").
		From(goqu.T("assignments")).
		Where(goqu.Ex{"user_id": userID}).
		ToSQL()
	if err != nil {
		return nil, ErrQueryBuildFailure
	}

	dest := []domain.AssignedTask{}
	if err := s.db.SelectContext(ctx, &dest, q, a...); err != nil {
		return nil, ErrDatabaseExecutionFailere
	}

	if len(dest) == 0 {
		return nil, nil
	}

	return &dest[0], nil
}

func (s *TaskStoreImpl) AssignTask(ctx context.Context, userID int, used *domain.TaskTopicUsed) (*domain.AssignedTask, error) {
	if s == nil {
		return nil, ErrNilReceiver
	}

	assigned, err := s.getAssignedTask(ctx, userID)
	if err != nil {
		return nil, err
	}

	if assigned != nil {
		return assigned, nil
	}

	b := dialect.
		Select("g.id", "g.task_id", "c.condition", "gc.counts").
		From(goqu.T("groups").As("g")).
		LeftJoin(
			goqu.T("conditions").As("c"),
			goqu.On(goqu.Ex{"g.condition_id": goqu.I("c.id")}),
		).
		LeftJoin(
			goqu.T("group_counts").As("gc"),
			goqu.On(goqu.Ex{"g.id": goqu.I("gc.group_id")}),
		)

	if used.Task1 && (used.Task1 != used.Task2) {
		b = b.Where(goqu.C("task_id").Eq(2))
	}

	if used.Task2 && (used.Task1 != used.Task2) {
		b = b.Where(goqu.C("task_id").Eq(1))
	}

	q, a, err := b.
		Order(goqu.I("gc.counts").Asc()).
		Limit(1).
		ToSQL()
	if err != nil {
		return nil, ErrQueryBuildFailure
	}

	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, ErrDatabaseExecutionFailere
	}

	dest := struct {
		ID        int    `db:"id"`
		TaskID    int    `db:"task_id"`
		Condition string `db:"condition"`
		Counts    int    `db:"counts"`
	}{}
	if err := tx.GetContext(ctx, &dest, q, a...); err != nil {
		return nil, ErrDatabaseExecutionFailere
	}

	q, a, err = dialect.
		Update("group_counts").
		Set(goqu.Record{"counts": dest.Counts + 1}).
		Where(goqu.Ex{"group_id": dest.ID}).
		ToSQL()

	if _, err := tx.ExecContext(ctx, q, a...); err != nil {
		return nil, ErrDatabaseExecutionFailere
	}

	q, a, err = dialect.
		Insert("assignments").
		Rows(goqu.Record{
			"user_id":   userID,
			"task_id":   dest.TaskID,
			"condition": dest.Condition,
		}).
		Prepared(true).
		ToSQL()
	if _, err := tx.ExecContext(ctx, q, a...); err != nil {
		return nil, ErrDatabaseExecutionFailere
	}

	tx.Commit()

	return &domain.AssignedTask{
		TaskID:    dest.TaskID,
		Condition: dest.Condition,
	}, nil
}

// CreateTask creates a new task record in DB
func (s *TaskStoreImpl) CreateTask(ctx context.Context, title, description string) (*domain.Task, error) {
	if s == nil {
		return nil, ErrNilReceiver
	}

	q, a, err := dialect.
		Insert("tasks").
		Rows(goqu.Record{"title": title, "description": description}).
		Prepared(true).
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

	return &domain.Task{
		ID:          id,
		Title:       title,
		Description: description,
	}, nil
}

// GetTaskSimple gets single instance of task for normal users.
func (s *TaskStoreImpl) GetTaskSimple(ctx context.Context, id int64) (*domain.TaskSimple, error) {
	if s == nil {
		return nil, ErrNilReceiver
	}

	q, a, err := dialect.Select("title", "description").From("tasks").ToSQL()
	if err != nil {
		return nil, ErrQueryBuildFailure
	}

	rs := &domain.TaskSimple{}
	if err := s.db.GetContext(ctx, rs, q, a...); err != nil {
		return nil, ErrDatabaseExecutionFailere
	}

	return rs, nil
}

// GetTask gets single instance of task for admin users.
func (s *TaskStoreImpl) GetTask(ctx context.Context, id int64) (*domain.Task, error) {
	if s == nil {
		return nil, ErrNilReceiver
	}

	q, a, err := dialect.
		Select("id", "topic", "query", "title", "description").
		From("tasks").
		Where(goqu.Ex{"id": id}).
		Prepared(true).
		ToSQL()
	if err != nil {
		return nil, ErrQueryBuildFailure
	}

	rs := &domain.Task{}
	if err := s.db.GetContext(ctx, rs, q, a...); err != nil {
		return nil, ErrDatabaseExecutionFailere
	}

	return rs, nil
}

// ListTasks gets a number of tasks for admin users.
func (s *TaskStoreImpl) ListTasks(ctx context.Context, limit int, offset int) ([]domain.Task, error) {
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

	rs := make([]domain.Task, limit)
	if err := s.db.SelectContext(ctx, rs, q, a...); err != nil {
		return nil, ErrDatabaseExecutionFailere
	}

	return rs, nil
}

// UpdateTasks updates a task for admin users.
func (s *TaskStoreImpl) UpdateTasks(ctx context.Context, task *domain.Task) (*domain.Task, error) {
	if s == nil {
		return nil, ErrNilReceiver
	}

	q, a, err := dialect.
		Update("tasks").
		Set(task).
		Prepared(true).
		ToSQL()
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

	q, a, err := dialect.
		Delete("tasks").
		Where(goqu.Ex{"id": id}).
		Prepared(true).
		ToSQL()
	if err != nil {
		return ErrQueryBuildFailure
	}

	if _, err := s.db.ExecContext(ctx, q, a...); err != nil {
		return ErrDatabaseExecutionFailere
	}

	return nil
}

func (s *TaskStoreImpl) CreateAnswer(ctx context.Context, answer *domain.Answer) error {
	if s == nil {
		return ErrNilReceiver
	}

	q, a, err := dialect.
		Insert("answers").
		Rows(goqu.Record{
			"user_id":   answer.UserID,
			"task_id":   answer.TaskID,
			"condition": answer.Condition,
			"answer":    answer.Answer,
			"reason":    answer.Reason,
		}).
		Prepared(true).
		ToSQL()
	if err != nil {
		return ErrQueryBuildFailure
	}

	if _, err := s.db.ExecContext(ctx, q, a...); err != nil {
		return ErrDatabaseExecutionFailere
	}

	return nil
}
