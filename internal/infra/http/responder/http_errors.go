package resp

import rest_err "github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"

func InvalidBodyRequest() error {
	return rest_err.NewBadRequestValidationError(
		"Invalid request body. The submitted data does not match the expected format.",
	)
}
