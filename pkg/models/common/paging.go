package common

type Paging struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}
