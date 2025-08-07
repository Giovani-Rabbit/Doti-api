package apperr

import "github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"

const (
	SttInternalServerError  = "INTERNAL_SERVER_ERROR"
	SttFieldValidationError = "FIELD_VALIDATION_ERROR"
)

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
