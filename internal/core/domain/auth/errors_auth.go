package authdomain

import rest_err "github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"

const (
	errorGeneratingToken = "ERROR_GENERATIONG_TOKEN"
)

func ErrGeneratingToken() *rest_err.RestErr {
	return rest_err.NewBadRequestError(
		errorGeneratingToken,
		"error generating token",
	)
}

func ErrGetUserFromContext() *rest_err.RestErr {
	return rest_err.NewBadRequestError(
		"INTERNAL_ERROR",
		"Error retrieving user data via token",
	)
}
