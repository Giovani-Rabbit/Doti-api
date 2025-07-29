package rest_err

import "net/http"

type RestErr struct {
	Message string `json:"message" example:"error trying to process request"`
	Status  string `json:"status" example:"USER_ALREADY_EXISTS"`
	Err     string `json:"error" example:"internal_server_error"`
	Code    int    `json:"code" example:"500"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewBadRequestError(status string, message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  status,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewUnauthorizedRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func NewBadRequestValidationError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Status:  "INVALID_REQUEST",
		Code:    http.StatusBadRequest,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
