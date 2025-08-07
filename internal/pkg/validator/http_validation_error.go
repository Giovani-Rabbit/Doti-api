package val

import "github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"

const (
	SttFieldValidationError = "FIELD_VALIDATION_ERROR"
	SttInternalServerError  = "INTERNAL_SERVER_ERROR"
	SttInvalidUUID          = "INVALID_UUID"
)

func ErrInvalidUUID() *http.RestErr {
	return http.NewBadRequestError(
		SttInvalidUUID,
		"id format invalid",
	)
}

func ErrValidationDomain(err error) *http.RestErr {
	return http.NewBadRequestError(
		SttFieldValidationError,
		err.Error(),
	)
}

func ErrInternal(msg string, err error) *http.RestErr {
	return http.NewRestError(
		SttInternalServerError,
		msg,
		err,
	)
}
