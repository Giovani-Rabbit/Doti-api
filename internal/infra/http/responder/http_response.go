package resp

import (
	"encoding/json"
	"net/http"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
)

type HttpResponse struct {
	writer http.ResponseWriter
	header map[string]string
	body   any
}

type IHttpResponse interface {
	AddHeader(key string, value string)
	AddBody(data any)
	DecodeJSONBody(r *http.Request, schema any) bool
	Error(err error)
	Write(statusCode int)
}

func NewHttpJSONResponse(w http.ResponseWriter) IHttpResponse {
	w.Header().Set("Content-Type", "application/json")

	return &HttpResponse{
		writer: w,
		header: map[string]string{},
		body:   nil,
	}
}

func (hr *HttpResponse) AddHeader(key string, value string) {
	hr.header[key] = value
}

func (hr *HttpResponse) AddBody(data any) {
	hr.body = data
}

func (hs *HttpResponse) DecodeJSONBody(r *http.Request, schema any) bool {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(schema); err != nil {
		logger.Error("Invalid request body.", nil)
		hs.Error(NewInvalidBodyRequest(err))
		return false
	}

	return true
}

func (hr *HttpResponse) Error(err error) {
	var resterr = AsRestErr(err)
	hr.AddBody(resterr)
	hr.Write(resterr.Code)
}

func (hr *HttpResponse) Write(statusCode int) {
	for k, v := range hr.header {
		hr.writer.Header().Add(k, v)
	}

	hr.writer.WriteHeader(statusCode)

	if hr.body != nil {
		json.NewEncoder(hr.writer).Encode(hr.body)
	}
}
