package types

type FilterPagination struct {
	TotalResults int       `json:"totalResults"`
	PageSize     int64      `json:"pageSize"`
	PageNumber   int64       `json:"pageNumber"`
	Results      interface{} `json:"results"`
}
