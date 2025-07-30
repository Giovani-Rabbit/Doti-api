package authdomain

import "github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"

const (
	errorGeneratingToken = "ERROR_GENERATIONG_TOKEN"
)

func ErrGeneratingToken() *http.RestErr {
	return http.NewBadRequestError(
		errorGeneratingToken,
		"error generating token",
	)
}

func ErrGetUserFromContext() *http.RestErr {
	return http.NewBadRequestError(
		"INTERNAL_ERROR",
		"Error retrieving user data via token",
	)
}
