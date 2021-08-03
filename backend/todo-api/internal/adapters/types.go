package adapters

type paginationMeta struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
