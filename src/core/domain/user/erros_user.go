package user

import rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"

const (
	userAlreadyExists = "USER_ALREADY_EXISTS"
	couldNotFindUser  = "COULD_NOT_FIND_USER"
)

func ErrUserAlreadyExists() *rest_err.RestErr {
	return rest_err.NewBadRequestError(
		userAlreadyExists,
		"A user with this email already exists.",
	)
}

func ErrCouldNotFindUser() *rest_err.RestErr {
	return rest_err.NewBadRequestError(
		couldNotFindUser,
		"Could not find user with those credentials",
	)
}
