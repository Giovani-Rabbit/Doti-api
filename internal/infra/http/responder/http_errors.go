package resp

import "net/http"

const (
	SttInvalidRequestBody = "INVALID_REQUEST_BODY"
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

func NewUnauthorizedRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
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
