package services

import (
	"blog-api/Domain/responses"
	"context"
	"log"
	"reflect"

	"github.com/mehdihadeli/go-mediatr"
)

func SendToMediator[TQuery any, TResponse any](ctx context.Context, query TQuery) TResponse {
	resp, err := mediatr.Send[TQuery, TResponse](ctx, query)
	if err != nil {
		log.Println(err)
		var response TResponse
		if ptr, ok := any(&response).(TResponse); ok {
			reflect.ValueOf(ptr).Elem().FieldByName("BaseResponse").
				Set(reflect.ValueOf(responses.NewBaseResponse(500, "Something went wrong")))
			return ptr
		}
		panic("TResponse must be a pointer type")
	}
	return resp
}
