package queries

type GetCarOrderQuery struct {
	ID string `json:"id" validate:"required"`
}
