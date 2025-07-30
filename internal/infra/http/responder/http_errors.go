package resp

import "github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"

func InvalidBodyRequest() error {
	return http.NewBadRequestValidationError(
		"Invalid request body. The submitted data does not match the expected format.",
	)
}
