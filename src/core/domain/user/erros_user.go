package userdomain

import rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"

const (
	userAlreadyExists      = "USER_ALREADY_EXISTS"
	couldNotFindUser       = "COULD_NOT_FIND_USER"
	userValuesMissing      = "USER_VALUES_MISSING"
	invalidUserEmailFormat = "INVALID_USER_EMAIL_FORMAT"
	invalidPassword        = "INVALID_PASSWORD"
)

func ErrUserAlreadyExists() error {
	return rest_err.NewBadRequestError(
		userAlreadyExists,
		"A user with this email already exists.",
	)
}

func ErrCouldNotFindUser() error {
	return rest_err.NewBadRequestError(
		couldNotFindUser,
		"Could not find user with those credentials",
	)
}

func ErrUserValuesMissing() error {
	return rest_err.NewBadRequestError(
		userValuesMissing,
		"User values are missing",
	)
}

func ErrSignInValuesMissing() error {
	return rest_err.NewBadRequestError(
		userValuesMissing,
		"Email or Password is missing",
	)
}

func ErrInvalidUserEmailFormat() error {
	return rest_err.NewBadRequestError(
		invalidUserEmailFormat,
		"Invalid user email format",
	)
}

func ErrInvalidPassword(message error) error {
	return rest_err.NewBadRequestError(
		invalidPassword,
		message.Error(),
	)
}
