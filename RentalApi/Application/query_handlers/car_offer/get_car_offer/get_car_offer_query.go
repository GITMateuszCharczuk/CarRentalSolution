package queries

type GetCarOfferQuery struct {
	ID string `json:"id" validate:"required"`
}
