package val

import "github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"

const (
	SttInvalidUUID = "INVALID_UUID"
)

func ErrInvalidUUID() *http.RestErr {
	return http.NewBadRequestError(
		SttInvalidUUID,
		"Invalid ID format",
	)
}
