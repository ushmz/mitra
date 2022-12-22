package store

import (
	"context"
	"fmt"
	"mitra/domain"

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
	CreateClickLog(context.Context, domain.ClickLog) error
	CreateDwelltimeLog(context.Context, domain.DwellTimeLog) error
	ListClickLogs(context.Context) ([]domain.ClickLog, error)
	ListDwellTimeLogs(context.Context) ([]domain.DwellTimeLog, error)
	GetClickLog(context.Context, GetLogsOpt) ([]domain.ClickLog, error)
	GetDwellTimeLog(context.Context, GetLogsOpt) ([]domain.DwellTimeLog, error)
}

type LogStoreImpl struct {
	db *sqlx.DB
}

// NewLogStore returns new Store object
func NewLogStore(db *sqlx.DB) LogStore {
	return &LogStoreImpl{db: db}
}

// CreateClickLog insert new row to DB
func (s *LogStoreImpl) CreateClickLog(ctx context.Context, p domain.ClickLog) error {
	if s == nil {
		return ErrNilReceiver
	}

	q, a, err := dialect.Insert("logs_event").Rows(p).ToSQL()
	if err != nil {
		return ErrQueryBuildFailure
	}

	if _, err := s.db.ExecContext(ctx, q, a...); err != nil {
		fmt.Println(err)
		return ErrDatabaseExecutionFailere
	}
	return nil
}

// CreateDwelltimeLog insert new row to DB
func (s *LogStoreImpl) CreateDwelltimeLog(ctx context.Context, p domain.DwellTimeLog) error {
	if s == nil {
		return ErrNilReceiver
	}

	q, a, err := dialect.
		Insert("logs_serp_dwell_time").
		Rows(p).
		OnConflict(
			goqu.DoUpdate(
				"time_on_page",
				goqu.Record{"time_on_page": goqu.L("time_on_page+1")})).
		ToSQL()
	if err != nil {
		return ErrQueryBuildFailure
	}

	if _, err := s.db.ExecContext(ctx, q, a...); err != nil {
		fmt.Println(err)
		return ErrDatabaseExecutionFailere
	}
	return nil
}

// ListClickLog lists all click logs
func (s *LogStoreImpl) ListClickLogs(ctx context.Context) ([]domain.ClickLog, error) {
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

	rs := []domain.ClickLog{}
	if err := s.db.Select(&rs, q, a...); err != nil {
		return nil, ErrDatabaseExecutionFailere
	}

	return rs, nil
}

// GetClickLogs gets specific click logs
func (s *LogStoreImpl) GetClickLog(ctx context.Context, opt GetLogsOpt) ([]domain.ClickLog, error) {
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

	rs := []domain.ClickLog{}
	if err := s.db.Select(&rs, q, a...); err != nil {
		return nil, ErrDatabaseExecutionFailere
	}

	return rs, nil
}

// ListDwellTimeLog lists all dwell time logs
func (s *LogStoreImpl) ListDwellTimeLogs(ctx context.Context) ([]domain.DwellTimeLog, error) {
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

	rs := []domain.DwellTimeLog{}
	if err := s.db.Select(&rs, q, a...); err != nil {
		return nil, ErrDatabaseExecutionFailere
	}

	return rs, nil
}

// GetDwellTimeLogs gets specific dwell time logs
func (s *LogStoreImpl) GetDwellTimeLog(ctx context.Context, opt GetLogsOpt) ([]domain.DwellTimeLog, error) {
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

	rs := []domain.DwellTimeLog{}
	if err := s.db.Select(&rs, q, a...); err != nil {
		return nil, ErrDatabaseExecutionFailere
	}

	return rs, nil
}
