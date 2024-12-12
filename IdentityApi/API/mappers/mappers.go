package mappers

import (
	"identity-api/API/services"
	get_all_users_contract "identity-api/Application.contract/get_all_users"
	get_user_info_contract "identity-api/Application.contract/get_user_info"
	get_user_internal_contract "identity-api/Application.contract/get_user_internal"
	login_contract "identity-api/Application.contract/login"
	modify_user_contract "identity-api/Application.contract/modify_user"
	register_contract "identity-api/Application.contract/register"
	validate_token_contract "identity-api/Application.contract/validate_token"
	login "identity-api/Application/command_handlers/login"
	modify_user "identity-api/Application/command_handlers/modify_user"
	register "identity-api/Application/command_handlers/register"
	validate_token "identity-api/Application/command_handlers/validate_token"
	get_all_users "identity-api/Application/query_handlers/get_all_users"
	get_user_info "identity-api/Application/query_handlers/get_user_info"
	get_user_internal "identity-api/Application/query_handlers/get_user_internal"
)

func MapToRegisterCommand(req *register_contract.RegisterUserRequest) register.RegisterUserCommand {
	return register.RegisterUserCommand{
		Name:         req.Name,
		Surname:      req.Surname,
		PhoneNumber:  req.PhoneNumber,
		EmailAddress: req.EmailAddress,
		Address:      req.Address,
		PostalCode:   req.PostalCode,
		City:         req.City,
		Password:     req.Password,
	}
}

func MapToLoginCommand(req *login_contract.LoginRequest) login.LoginCommand {
	return login.LoginCommand{
		Email:    req.Email,
		Password: req.Password,
	}
}

func MapToModifyUserCommand(req *modify_user_contract.ModifyUserRequest) modify_user.ModifyUserCommand {
	return modify_user.ModifyUserCommand{
		UserID:      req.UserID,
		Name:        req.Name,
		Surname:     req.Surname,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		PostalCode:  req.PostalCode,
		City:        req.City,
	}
}

func MapToValidateTokenCommand(req *validate_token_contract.ValidateTokenRequest) validate_token.ValidateTokenCommand {
	return validate_token.ValidateTokenCommand{
		JwtToken: req.JwtToken,
	}
}

func MapToGetAllUsersQuery(req *get_all_users_contract.GetAllUsersRequest) get_all_users.GetAllUsersQuery {
	return get_all_users.GetAllUsersQuery{
		JwtToken:   req.JwtToken,
		Pagination: req.Pagination,
		Sortable:   services.ExtractSorting(req.SortQuery),
	}
}

func MapToGetUserInternalQuery(req *get_user_internal_contract.GetUserInternalRequest) get_user_internal.GetUserInternalQuery {
	return get_user_internal.GetUserInternalQuery{
		JwtToken: req.JwtToken,
	}
}

func MapToGetUserInfoQuery(req *get_user_info_contract.GetUserInfoRequest) get_user_info.GetUserInfoQuery {
	return get_user_info.GetUserInfoQuery{
		JwtToken: req.JwtToken,
	}
}
