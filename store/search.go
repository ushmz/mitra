package store

import (
	"fmt"
	"mitra/domain"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type SearchStore interface {
	GetSearchResults(ctx context.Context, offset, limit int) ([]domain.SearchResult, error)
	GetSearchPages(ctx context.Context, taskID, offset, limit, top int) ([]domain.Result, error)
	GetSimilarwebPagesByPageIDs(ctx context.Context, taskIDs []int) ([]interface{}, error)
}

type SearchStoreImpl struct {
	db *sqlx.DB
}

func NewSearchStore(db *sqlx.DB) SearchStore {
	return &SearchStoreImpl{db: db}
}

func (s *SearchStoreImpl) GetSearchResults(ctx context.Context, offset, limit int) ([]domain.SearchResult, error) {
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

func (s *SearchStoreImpl) GetSearchPages(ctx context.Context, taskID, offset, limit, top int) ([]domain.Result, error) {
	if s == nil {
		return nil, ErrNilReceiver
	}

	if offset < 0 {
		offset = 0
	}

	if offset > 3 {
		offset = 3
	}

	if limit < 0 || limit > 10 {
		limit = 10
	}

	q, a, err := dialect.
		From(dialect.
			From(goqu.T("search_pages").As("s")).
			Select("*").
			Where(goqu.C("task_id").Eq(taskID)).
			Offset(uint(offset*limit)).
			Limit(uint(limit))).
		Select(
			goqu.C("pages.id").As("page_id"),
			goqu.C("pages.title").As("page_title"),
			goqu.C("pages.url").As("page_url"),
			goqu.C("pages.snippet").As("page_snippet"),
			goqu.C("sp.title").As("similarweb_title"),
			goqu.C("sp.url").As("similarweb_url"),
			goqu.C("sp.icon_path").As("similarweb_icon_path"),
			goqu.C("c.category").As("similarweb_category")).
		LeftJoin(
			dialect.
				Select(
					"r.page_id",
					"r.task_id",
					"r.similarweb_id",
					"r.idf",
					goqu.ROW_NUMBER().
						Over(goqu.W().
							PartitionBy("page_id").
							OrderBy(goqu.C("idf").Desc())).
						As("idf_rank"),
				).
				From(goqu.T("search_page_wimilarweb_relations").As("r")).
				As("idf_ranks"),
			goqu.On(
				goqu.Ex{"idf_ranks.page_id": goqu.I("pages.id")},
				goqu.C("idf_ranks.idf_rank").Lte(10),
			)).
		LeftJoin(
			goqu.T("idf_ranks").As("i"),
			goqu.On(
				goqu.Ex{"i.page_id": goqu.I("pages.id")},
				goqu.C("i.idf_rank").Lte(10))).
		LeftJoin(
			goqu.T("similarweb_pages").As("sp"),
			goqu.On(goqu.Ex{"sp.id": goqu.I("idf_ranks.similarweb_id")})).
		LeftJoin(
			goqu.T("similarweb_categories").As("c"),
			goqu.On(goqu.Ex{"sp.category": goqu.I("c.category")})).
		Where(goqu.C("i.idf_rank").Lte(10)).
		ToSQL()
	if err != nil {
		fmt.Println("Query: ", err)
		return nil, ErrQueryBuildFailure
	}

	rs := []domain.RelationResult{}
	if err := s.db.SelectContext(ctx, &rs, q, a...); err != nil {
		fmt.Println(q, a)
		fmt.Println("Exec: ", err)
		return nil, ErrDatabaseExecutionFailere
	}

	resultMap := map[int]domain.Result{}
	for _, r := range rs {
		if _, ok := resultMap[r.ID]; !ok {
			resultMap[r.ID] = domain.Result{
				ID:              r.ID,
				Title:           r.Title,
				URL:             r.URL,
				Snippet:         r.Snippet,
				SimilarwebPages: []domain.SimilarwebPage{},
			}

			temp := resultMap[r.ID]
			tmp := temp.SimilarwebPages
			tmp = append(tmp, domain.SimilarwebPage{
				ID:       r.SimilarwebID,
				Title:    r.SimilarwebTitle,
				URL:      r.SimilarwebURL,
				Icon:     r.SimilarwebIcon,
				Category: r.SimilarwebCategory,
			})
			temp.SimilarwebPages = tmp
			resultMap[r.ID] = temp
		}
	}

	result := []domain.Result{}
	for _, v := range resultMap {
		result = append(result, v)
	}

	return result, nil
}

func (s *SearchStoreImpl) GetSimilarwebPagesByPageIDs(ctx context.Context, taskIDs []int) ([]interface{}, error) {
	if s == nil {
		return nil, ErrNilReceiver
	}

	q, a, err := dialect.
		Select(
			goqu.C("sp.page_id").As("page_id"),
			goqu.C("sp.title").As("similarweb_title"),
			goqu.C("sp.url").As("similarweb_url"),
			goqu.C("sp.icon_path").As("similarweb_icon_path"),
			goqu.C("c.category").As("similarweb_category")).
		From(dialect.
			Select(
				goqu.C("r.page_id").As("page_id"),
				goqu.C("r.task_id").As("task_id"),
				goqu.C("r.similarweb_id").As("similarweb_id"),
				goqu.C("r.idf").As("idf"),
				goqu.ROW_NUMBER().
					Over(goqu.W().
						PartitionBy("page_id")).
					As("idf_rank"),
			).
			From(goqu.T("search_page_wimilarweb_relations").As("r")).
			Where(goqu.C("r.task_id").In(taskIDs)).
			As("i")).
		LeftJoin(
			goqu.T("similarweb_pages").As("sp"),
			goqu.On(goqu.Ex{"sp.id": goqu.I("idf_ranks.similarweb_id")})).
		LeftJoin(
			goqu.T("similarweb_categories").As("c"),
			goqu.On(goqu.Ex{"sp.category": goqu.I("c.category")})).
		Where(goqu.C("i.idf_rank").Lte(10)).
		ToSQL()
	if err != nil {
		fmt.Println(err)
		return nil, ErrQueryBuildFailure
	}

	fmt.Println(q, a)

	return nil, nil
}
