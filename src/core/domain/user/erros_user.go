package userDomain

import rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"

const (
	userAlreadyExists = "USER_ALREADY_EXISTS"
)

func ErrUserAlreadyExists() *rest_err.RestErr {
	return rest_err.NewBadRequestError(
		userAlreadyExists,
		"A user with this email already exists.",
	)
}
