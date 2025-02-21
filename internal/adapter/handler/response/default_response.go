package response

type ErrorResponseDefault struct {
	Meta
}

type Meta struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type DefaultSuccessResponse struct {
	Meta       Meta                `json:"meta"`
	Data       interface{}         `json:"data,omitempty"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

type PaginationResponse struct {
	TotalRecords int `json:"total_records"`
	Page         int `json:"page"`
	PerPage      int `json:"per_page"`
	TotalPages   int `json:"total_pages"`
}
