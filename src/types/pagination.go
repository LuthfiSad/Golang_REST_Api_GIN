package types

type PaginationMeta struct {
	PerPage    int `json:"perPage"`
	Page       int `json:"page"`
	TotalData  int `json:"totalData"`
	TotalPages int `json:"totalPages"`
}
