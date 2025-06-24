package user

import (
	"context"
	"encoding/json"
	"net/http"

	userdto "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/user/dtos"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
)

func (uc *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Invalid Content-Type", http.StatusBadRequest)
		return
	}

	var user userdto.CreateUserDTO
	if err := decodeJSONBody(w, r, &user); err != nil {
		handleError(w, err)
		return
	}

	ctx := context.Background()

	res, err := uc.CreateUserUseCase.Execute(ctx, user)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userdto.NewUserCreatedResponse(res))
}

func decodeJSONBody(_ http.ResponseWriter, r *http.Request, dst interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		return rest_err.NewBadRequestError("UNGW", "Invalid JSON body")
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
