package pagination

type Pagination struct {
	Page     int `form:"page" json:"page" example:"1"`
	PageSize int `form:"page_size" json:"page_size" example:"10"`
}
