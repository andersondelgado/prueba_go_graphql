package global

type PaginationSimpleParams struct {
	Filter  string `json:"filter"`
	Offset  int    `json:"offset"`
	Limit   int    `json:"limit"`
	OrderBy string `json:"orderBy"`
}