package userdomain

import "github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"

const (
	SttCouldNotFindUser       = "COULD_NOT_FIND_USER"
	SttCouldNotPersistUser    = "COULD_NOT_PERSIST_USER"
	SttInvalidUserEmailFormat = "INVALID_USER_EMAIL_FORMAT"
	SttInvalidPassword        = "INVALID_PASSWORD"
	SttErrorGeneratingToken   = "ERROR_GENERATIONG_TOKEN"
	SttUserValuesMissing      = "USER_VALUES_MISSING"
	SttUserAlreadyExists      = "USER_ALREADY_EXISTS"
)

func ErrGeneratingToken() *http.RestErr {
	return http.NewBadRequestError(
		SttErrorGeneratingToken,
		"error generating token",
	)
}

func ErrGetUserFromContext() *http.RestErr {
	return http.NewBadRequestError(
		"INTERNAL_ERROR",
		"Error retrieving user data via token",
	)
}

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
		"Missing required fields. Please check that all required fields are provided.",
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

func ErrCouldNotPersistUser(err error) *http.RestErr {
	return http.ErrInternal(
		"Error saving user",
		SttCouldNotPersistUser,
		err,
	)
}
