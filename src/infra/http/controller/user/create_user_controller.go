package userController

import (
	"context"
	"encoding/json"
	"net/http"

	userDTO "github.com/Giovani-Coelho/Doti-API/src/infra/http/controller/user/dtos"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
)

func (uc *UserControllers) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user userDTO.CreateUserDTO

	if err := decodeJSONBody(w, r, &user); err != nil {
		handleError(w, err)
		return
	}

	ctx := context.Background()
	if err := uc.UserServices.CreateUser(ctx, user); err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Account Created"))
}

func decodeJSONBody(_ http.ResponseWriter, r *http.Request, dst interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		return rest_err.NewBadRequestError("Invalid JSON body")
	}
	return nil
}

func handleError(w http.ResponseWriter, err error) {
	httpErr, ok := err.(*rest_err.RestErr)
	if !ok {
		httpErr = rest_err.NewInternalServerError("Unexpected error" + err.Error())
	}

	response, _ := json.Marshal(httpErr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpErr.Code)
	w.Write(response)
}
