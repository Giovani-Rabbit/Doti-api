package httphdl

import (
	"encoding/json"
	"net/http"

	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
)

func DecodeJSONBody(r *http.Request, dst interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(dst); err != nil {
		return err
	}

	return nil
}

func HandleError(w http.ResponseWriter, err error) {
	httpErr, ok := err.(*rest_err.RestErr)

	if !ok {
		httpErr = rest_err.NewInternalServerError("Unexpected error" + err.Error())
	}

	response, _ := json.Marshal(httpErr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpErr.Code)
	w.Write(response)
}

func ResponseHttpJson(w http.ResponseWriter, status int, jsonResponse any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(jsonResponse)
}
