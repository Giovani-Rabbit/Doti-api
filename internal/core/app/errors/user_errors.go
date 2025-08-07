package apperr

import "github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"

const (
	SttUserAlreadyExists      = "USER_ALREADY_EXISTS"
	SttCouldNotFindUser       = "COULD_NOT_FIND_USER"
	SttUserValuesMissing      = "USER_VALUES_MISSING"
	SttInvalidUserEmailFormat = "INVALID_USER_EMAIL_FORMAT"
	SttInvalidPassword        = "INVALID_PASSWORD"
)

func ErrUserAlreadyExists() *http.RestErr {
	return http.NewBadRequestError(
		SttUserAlreadyExists,
		"A user with this email already exists.",
	)
}

func ErrCouldNotFindUser() *http.RestErr {
	return http.NewBadRequestError(
		SttCouldNotFindUser,
		"Could not find user with those credentials",
	)
}

func ErrUserValuesMissing() *http.RestErr {
	return http.NewBadRequestError(
		SttUserValuesMissing,
		"User values are missing",
	)
}

func ErrSignInValuesMissing() *http.RestErr {
	return http.NewBadRequestError(
		SttUserValuesMissing,
		"Email or Password is missing",
	)
}

func ErrInvalidUserEmailFormat() *http.RestErr {
	return http.NewBadRequestError(
		SttInvalidUserEmailFormat,
		"Invalid user email format",
	)
}

func ErrInvalidPassword(message error) *http.RestErr {
	return http.NewBadRequestError(
		SttInvalidPassword,
		message.Error(),
	)
}
