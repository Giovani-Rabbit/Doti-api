package userdomain

import rest_err "github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"

const (
	SttUserAlreadyExists      = "USER_ALREADY_EXISTS"
	SttCouldNotFindUser       = "COULD_NOT_FIND_USER"
	SttUserValuesMissing      = "USER_VALUES_MISSING"
	SttInvalidUserEmailFormat = "INVALID_USER_EMAIL_FORMAT"
	SttInvalidPassword        = "INVALID_PASSWORD"
)

func ErrUserAlreadyExists() *rest_err.RestErr {
	return rest_err.NewBadRequestError(
		SttUserAlreadyExists,
		"A user with this email already exists.",
	)
}

func ErrCouldNotFindUser() *rest_err.RestErr {
	return rest_err.NewBadRequestError(
		SttCouldNotFindUser,
		"Could not find user with those credentials",
	)
}

func ErrUserValuesMissing() *rest_err.RestErr {
	return rest_err.NewBadRequestError(
		SttUserValuesMissing,
		"User values are missing",
	)
}

func ErrSignInValuesMissing() *rest_err.RestErr {
	return rest_err.NewBadRequestError(
		SttUserValuesMissing,
		"Email or Password is missing",
	)
}

func ErrInvalidUserEmailFormat() *rest_err.RestErr {
	return rest_err.NewBadRequestError(
		SttInvalidUserEmailFormat,
		"Invalid user email format",
	)
}

func ErrInvalidPassword(message error) *rest_err.RestErr {
	return rest_err.NewBadRequestError(
		SttInvalidPassword,
		message.Error(),
	)
}
