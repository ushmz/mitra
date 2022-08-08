package store

import (
	"context"
	"mitra/domain/model"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

// GetLogsOpt : Log filtering options
type GetLogsOpt struct {
	UserID      int
	TaskID      int
	ConditionID int
}

type LogStore interface {
	CreateClickLog(context.Context, model.ClickLog) error
	CreateDwelltimeLog(context.Context, model.DwellTimeLog) error
	ListClickLog(context.Context) ([]model.ClickLog, error)
	ListDwellTimeLog(context.Context) ([]model.DwellTimeLog, error)
	GetClickLogs(context.Context, GetLogsOpt) ([]model.ClickLog, error)
	GetDwellTimeLogs(context.Context, GetLogsOpt) ([]model.DwellTimeLog, error)
}

type LogStoreImpl struct {
	db *sqlx.DB
}

// NewLogStore returns new Store object
func NewLogStore(db *sqlx.DB) LogStore {
	return &LogStoreImpl{db: db}
}

// CreateClickLog insert new row to DB
func (s *LogStoreImpl) CreateClickLog(ctx context.Context, p model.ClickLog) error {
	if s == nil {
		return ErrNilReceiver
	}

	q, a, err := dialect.Insert("click_logs").Rows(p).ToSQL()
	if err != nil {
		return ErrQueryBuildFailure
	}

	if _, err := s.db.ExecContext(ctx, q, a); err != nil {
		return ErrDatabaseExecutionFailere
	}
	return nil
}

// CreateDwelltimeLog insert new row to DB
func (s *LogStoreImpl) CreateDwelltimeLog(ctx context.Context, p model.DwellTimeLog) error {
	if s == nil {
		return ErrNilReceiver
	}

	q, a, err := dialect.Insert("dwell_time_logs").Rows(p).ToSQL()
	if err != nil {
		return ErrQueryBuildFailure
	}

	if _, err := s.db.ExecContext(ctx, q, a); err != nil {
		return ErrDatabaseExecutionFailere
	}
	return nil
}

// ListClickLog lists all click logs
func (s *LogStoreImpl) ListClickLog(ctx context.Context) ([]model.ClickLog, error) {
	if s == nil {
		return nil, ErrNilReceiver
	}

	q, a, err := dialect.
		Select("*").
		From("logs_event").
		ToSQL()
	if err != nil {
		return nil, ErrQueryBuildFailure
	}

	rs := []model.ClickLog{}
	if err := s.db.Select(&rs, q, a...); err != nil {
		return nil, ErrDatabaseExecutionFailere
	}

	return rs, nil
}

// GetClickLogs gets specific click logs
func (s *LogStoreImpl) GetClickLogs(ctx context.Context, opt GetLogsOpt) ([]model.ClickLog, error) {
	if s == nil {
		return nil, ErrNilReceiver
	}

	b := dialect.
		Select("*").
		From("logs_event")

	if opt.UserID > 0 {
		b = b.Where(goqu.Ex{"user_id": opt.UserID})
	}

	if opt.TaskID > 0 {
		b = b.Where(goqu.Ex{"task_id": opt.TaskID})
	}

	if opt.ConditionID > 0 {
		b = b.Where(goqu.Ex{"condition_id": opt.ConditionID})
	}

	q, a, err := b.ToSQL()
	if err != nil {
		return nil, ErrQueryBuildFailure
	}

	rs := []model.ClickLog{}
	if err := s.db.Select(&rs, q, a...); err != nil {
		return nil, ErrDatabaseExecutionFailere
	}

	return rs, nil
}

// ListDwellTimeLog lists all dwell time logs
func (s *LogStoreImpl) ListDwellTimeLog(ctx context.Context) ([]model.DwellTimeLog, error) {
	if s == nil {
		return nil, ErrNilReceiver
	}

	q, a, err := dialect.
		Select("*").
		From("logs_serp_dwell_time").
		ToSQL()
	if err != nil {
		return nil, ErrQueryBuildFailure
	}

	rs := []model.DwellTimeLog{}
	if err := s.db.Select(&rs, q, a...); err != nil {
		return nil, ErrDatabaseExecutionFailere
	}

	return rs, nil
}

// GetDwellTimeLogs gets specific dwell time logs
func (s *LogStoreImpl) GetDwellTimeLogs(ctx context.Context, opt GetLogsOpt) ([]model.DwellTimeLog, error) {
	if s == nil {
		return nil, ErrNilReceiver
	}

	b := dialect.
		Select("*").
		From("logs_serp_dwell_time")

	if opt.UserID > 0 {
		b = b.Where(goqu.Ex{"user_id": opt.UserID})
	}

	if opt.TaskID > 0 {
		b = b.Where(goqu.Ex{"task_id": opt.TaskID})
	}

	if opt.ConditionID > 0 {
		b = b.Where(goqu.Ex{"condition_id": opt.ConditionID})
	}

	q, a, err := b.ToSQL()
	if err != nil {
		return nil, ErrQueryBuildFailure
	}

	rs := []model.DwellTimeLog{}
	if err := s.db.Select(&rs, q, a...); err != nil {
		return nil, ErrDatabaseExecutionFailere
	}

	return rs, nil
}
