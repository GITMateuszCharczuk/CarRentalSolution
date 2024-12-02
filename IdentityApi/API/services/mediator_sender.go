package services

import (
	"context"
	"identity-api/Domain/responses"

	"github.com/mehdihadeli/go-mediatr"
)

func SendToMediator[TQuery any, TResponse any](ctx context.Context, query TQuery) TResponse {
	resp, err := mediatr.Send[TQuery, TResponse](ctx, query)
	if err != nil {
		response := responses.NewBaseResponse(500, "Something went wrong")
		return any(response).(TResponse)
	}
	return resp
}
