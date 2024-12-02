package requests

type Pagination struct {
	Page     int `json:"page" form:"page" binding:"required" example:"1" validate:"required,min=1"`
	PageSize int `json:"page_size" form:"page_size" binding:"required" example:"10" validate:"required,min=1,max=100"`
}
