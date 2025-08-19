package resp

import "github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"

const SttInvalidRequestBody = "INVALID_REQUEST_BODY"

func InvalidBodyRequest(parseErr error) error {
	return &http.RestErr{
		Message: "Invalid request body. The submitted data does not match the expected format.",
		Status:  SttInvalidRequestBody,
		Err:     parseErr.Error(),
		Code:    400,
	}
}
