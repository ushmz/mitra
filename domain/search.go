package domain

type SearchResult struct {
	ID      int    `db:"id" json:"-"`
	Title   string `db:"title" json:"title"`
	URL     string `db:"url" json:"url"`
	Snippet string `db:"snippet" json:"snippet"`
}

type SimilarwebPage struct {
	ID       int    `db:"id" json:"-"`
	Title    string `db:"title" json:"title"`
	URL      string `db:"url" json:"url"`
	Icon     string `db:"icon_path" json:"icon"`
	Category string `db:"category" json:"category"`
}

type RelationResult struct {
	ID                 int    `db:"page_id"`
	Title              string `db:"page_title"`
	URL                string `db:"page_url"`
	Snippet            string `db:"page_snippet"`
	SimilarwebID       int    `db:"similarweb_id"`
	SimilarwebTitle    string `db:"similarweb_title"`
	SimilarwebURL      string `db:"similarweb_url"`
	SimilarwebIcon     string `db:"similarweb_icon_path"`
	SimilarwebCategory string `db:"similarweb_category"`
}

type Result struct {
	ID              int              `json:"id"`
	Title           string           `json:"title"`
	URL             string           `json:"url"`
	Snippet         string           `json:"snippet"`
	SimilarwebPages []SimilarwebPage `json:"similarweb"`
}
