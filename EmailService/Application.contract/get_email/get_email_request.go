package contract

type GetEmailRequest struct {
	ID string `json:"id" binding:"required" validate:"required,len=60"`
}
