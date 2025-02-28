package response

type ServiceDetailResponse struct {
	ID          int64   `json:"id"`
	ServiceID   int64   `json:"service_id"`
	PathImage   string  `json:"path_image"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	PathPdf     *string `json:"path_pdf"`
	PathDocx    *string `json:"path_docx"`
	ServiceName string  `json:"service_name"`
}
