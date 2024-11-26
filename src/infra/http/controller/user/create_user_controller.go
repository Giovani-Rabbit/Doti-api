package userController

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	userDTO "github.com/Giovani-Coelho/Doti-API/src/application/user/dtos"
	userServices "github.com/Giovani-Coelho/Doti-API/src/application/user/services/createUser"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
)

type CreateUserController struct {
	UserServices userServices.ICreateUserService
}

type ICreateUserController interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

func NewCreateUserController(
	userServices userServices.ICreateUserService,
) ICreateUserController {
	return &CreateUserController{
		UserServices: userServices,
	}
}

func (s *CreateUserController) Execute(w http.ResponseWriter, r *http.Request) {
	var user userDTO.CreateUserDTO

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		httpErr := rest_err.NewBadRequestError("Unable to parse request body")

		res, err := json.Marshal(httpErr)
		if err != nil {
			log.Fatal(err)
		}

		WriteJSON(w, res)
		return
	}

	ctx := context.Background()

	err = s.UserServices.CreateUser(ctx, user)
	if err != nil {
		if httpErr, ok := err.(*rest_err.RestErr); ok {
			res, err := json.Marshal(httpErr)
			if err != nil {
				log.Fatal(err)
			}

			WriteJSON(w, res)
			return
		}

		return
	}
}

func WriteJSON(w http.ResponseWriter, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)
	w.Write(response)
}
