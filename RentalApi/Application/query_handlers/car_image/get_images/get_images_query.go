package queries

type GetImagesQuery struct {
	CarOfferId string `json:"carOfferId" validate:"required"`
}
