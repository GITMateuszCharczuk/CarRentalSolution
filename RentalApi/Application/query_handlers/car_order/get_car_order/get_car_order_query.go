package queries

import models "rental-api/Domain/models/external"

type GetCarOrderQuery struct {
	ID              string `json:"id" validate:"required"`
	models.JwtToken `json:",inline"`
}
