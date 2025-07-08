package authdomain

import rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"

const (
	errorGeneratingToken = "ERROR_GENERATIONG_TOKEN"
)

func ErrGeneratingToken() error {
	return rest_err.NewBadRequestError(
		errorGeneratingToken,
		"error generating token",
	)
}
