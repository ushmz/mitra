package store

import (
	"mitra/domain"

	"github.com/jmoiron/sqlx"
)

type SearchStore interface {
	GetSearchResults(offset, limit int) ([]domain.SearchResult, error)
}

type SearchStoreImpl struct {
	db *sqlx.DB
}

func NewSearchStore(db *sqlx.DB) SearchStore {
	return &SearchStoreImpl{db: db}
}

func (s *SearchStoreImpl) GetSearchResults(offset, limit int) ([]domain.SearchResult, error) {
	if s == nil {
		return nil, ErrNilReceiver
	}

	q, a, err := dialect.
		Select("*").
		From("search_pages").
		Limit(uint(limit)).
		Offset(uint(offset * limit)).
		ToSQL()
	if err != nil {
		return nil, ErrQueryBuildFailure
	}

	rs := make([]domain.SearchResult, limit)
	if err := s.db.Select(&rs, q, a...); err != nil {
		return nil, ErrDatabaseExecutionFailere
	}

	return rs, nil
}
