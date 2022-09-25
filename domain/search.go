package domain

type SearchResult struct {
	ID      int    `db:"id" json:"-"`
	Title   string `db:"title" json:"title"`
	URL     string `db:"url" json:"url"`
	Snippet string `db:"snippet" json:"snippet"`
}
