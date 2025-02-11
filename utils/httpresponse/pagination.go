package httpresponse

type Pagination struct {
	Items      []interface{} `json:"items"`
	TotalItems int           `json:"totalItems"`
	TotalPages int           `json:"totalPages"`
}

type ResponsesPagination struct {
	Items      interface{} `json:"items"`
	TotalItems int         `json:"totalItems"`
	TotalPages int         `json:"totalPages"`
}
