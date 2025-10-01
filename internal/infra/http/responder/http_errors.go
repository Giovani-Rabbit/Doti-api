package resp

import (
	"errors"
	"net/http"
)

const (
	SttInvalidRequestBody  = "INVALID_REQUEST_BODY"
	SttInvalidErrorType    = "INVALID_ERROR_TYPE"
	SttInvalidPathValue    = "INVALID_PATH_VALUE"
	SttInternalServerError = "INTERNAL_SERVER_ERROR"
)

type RestErr struct {
	Message string `json:"message" example:"error trying to process request"`
	Status  string `json:"status" example:"USER_ALREADY_EXISTS"`
	Err     string `json:"error" example:"internal_server_error"`
	Code    int    `json:"code" example:"500"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func AsRestErr(e error) *RestErr {
	if e == nil {
		return &RestErr{
			Message: "AsRestErr called with nil error",
			Status:  SttInternalServerError,
			Err:     "invalid error type",
			Code:    http.StatusInternalServerError,
		}
	}

	var restErr *RestErr
	if errors.As(e, &restErr) {
		return restErr
	}

	return &RestErr{
		Message: "unexpected error occurred",
		Status:  SttInternalServerError,
		Err:     "error of type RestErr expected",
		Code:    http.StatusInternalServerError,
	}
}

func NewBadRequestError(status string, message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  status,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewErrInternal(msg string, stt string, err error) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  stt,
		Err:     err.Error(),
		Code:    http.StatusInternalServerError,
	}
}

func NewUnauthorizedRequestError(status string, message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  status,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func NewCouldNotFind(message string, status string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  status,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewInvalidBodyRequest(parseErr error) error {
	return &RestErr{
		Message: "Invalid request body. The submitted data does not match the expected format.",
		Status:  SttInvalidRequestBody,
		Err:     parseErr.Error(),
		Code:    http.StatusBadRequest,
	}
}

func NewInvalidPathValue(err error) error {
	return &RestErr{
		Message: "Invalid path value",
		Status:  SttInvalidPathValue,
		Err:     err.Error(),
		Code:    http.StatusBadRequest,
	}
}
