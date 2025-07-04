package httphdl

import (
	"encoding/json"
	"net/http"

	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
)

type HttpResponse struct {
	writer http.ResponseWriter
	header map[string]string
	body   any
}

type IHttpResponse interface {
	AddHeader(string, string)
	AddBody(any)
	Error()
	SetStatusCode(int)
	Write()
}

func NewHttpJSONResponse(w http.ResponseWriter) *HttpResponse {
	w.Header().Set("Content-Type", "application/json")

	return &HttpResponse{
		writer: w,
		header: map[string]string{},
	}
}

func (hr *HttpResponse) AddHeader(key string, value string) {
	hr.header[key] = value
}

func (hr *HttpResponse) AddBody(data any) {
	hr.body = data
}

func (hr *HttpResponse) Write(code int) error {
	for k, v := range hr.header {
		hr.writer.Header().Add(k, v)
	}

	hr.writer.WriteHeader(code)

	return json.NewEncoder(hr.writer).Encode(hr.body)
}

func (hr *HttpResponse) Error(err error, code int) {
	httpErr := rest_err.NewInternalServerError(
		"Unexpected error" + err.Error(),
	)

	hr.AddBody(httpErr)
	hr.Write(code)
}

func DecodeJSONBody(r *http.Request, dst interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(dst); err != nil {
		return err
	}

	return nil
}
