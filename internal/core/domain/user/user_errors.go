package userdomain

import resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"

const (
	SttCouldNotFindUser       = "COULD_NOT_FIND_USER"
	SttCouldNotPersistUser    = "COULD_NOT_PERSIST_USER"
	SttInvalidUserEmailFormat = "INVALID_USER_EMAIL_FORMAT"
	SttInvalidPassword        = "INVALID_PASSWORD"
	SttErrorGeneratingToken   = "ERROR_GENERATIONG_TOKEN"
	SttUserValuesMissing      = "USER_VALUES_MISSING"
	SttUserAlreadyExists      = "USER_ALREADY_EXISTS"
)

func ErrGeneratingToken() *resp.RestErr {
	return resp.NewBadRequestError(
		SttErrorGeneratingToken,
		"error generating token",
	)
}

func ErrGetUserFromContext() *resp.RestErr {
	return resp.NewBadRequestError(
		"INTERNAL_ERROR",
		"Error retrieving user data via token",
	)
}

func ErrUserAlreadyExists() *resp.RestErr {
	return resp.NewBadRequestError(
		SttUserAlreadyExists,
		"A user with this email already exists.",
	)
}

func ErrCouldNotFindUser() *resp.RestErr {
	return resp.NewBadRequestError(
		SttCouldNotFindUser,
		"Could not find user with those credentials",
	)
}

func ErrUserValuesMissing() *resp.RestErr {
	return resp.NewBadRequestError(
		SttUserValuesMissing,
		"Missing required fields. Please check that all required fields are provided.",
	)
}

func ErrSignInValuesMissing() *resp.RestErr {
	return resp.NewBadRequestError(
		SttUserValuesMissing,
		"Email or Password is missing",
	)
}

func ErrInvalidUserEmailFormat() *resp.RestErr {
	return resp.NewBadRequestError(
		SttInvalidUserEmailFormat,
		"Invalid user email format",
	)
}

func ErrInvalidPassword(message error) *resp.RestErr {
	return resp.NewBadRequestError(
		SttInvalidPassword,
		message.Error(),
	)
}

func ErrCouldNotPersistUser(err error) *resp.RestErr {
	return resp.NewErrInternal(
		"Error saving user",
		SttCouldNotPersistUser,
		err,
	)
}
