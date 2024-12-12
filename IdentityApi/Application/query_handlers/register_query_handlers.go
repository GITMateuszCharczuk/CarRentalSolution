package queries

import (
	get_all_users "identity-api/Application/query_handlers/get_all_users"
	get_user_info "identity-api/Application/query_handlers/get_user_info"
	get_user_internal "identity-api/Application/query_handlers/get_user_internal"
	repository_interfaces "identity-api/Domain/repository_interfaces/user_repository"
	service_interfaces "identity-api/Domain/service_interfaces"
	"log"

	"github.com/mehdihadeli/go-mediatr"
)

func registerGetAllUsersQueryHandler(
	userQueryRepository repository_interfaces.UserQueryRepository,
	tokenService service_interfaces.JWTTokenService,
) {
	handler := get_all_users.NewGetAllUsersQueryHandler(userQueryRepository, tokenService)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerGetUserInternalQueryHandler(
	userQueryRepository repository_interfaces.UserQueryRepository,
	tokenService service_interfaces.JWTTokenService,
) {
	handler := get_user_internal.NewGetUserInternalQueryHandler(userQueryRepository, tokenService)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerGetUserInfoQueryHandler(
	userQueryRepository repository_interfaces.UserQueryRepository,
	tokenService service_interfaces.JWTTokenService,
) {
	handler := get_user_info.NewGetUserInfoQueryHandler(userQueryRepository, tokenService)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func RegisterQueryHandlers(
	userQueryRepository repository_interfaces.UserQueryRepository,
	tokenService service_interfaces.JWTTokenService,
) {
	registerGetAllUsersQueryHandler(userQueryRepository, tokenService)
	registerGetUserInternalQueryHandler(userQueryRepository, tokenService)
	registerGetUserInfoQueryHandler(userQueryRepository, tokenService)
}
