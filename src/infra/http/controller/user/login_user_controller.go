package userController

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	authDTO "github.com/Giovani-Coelho/Doti-API/src/application/auth/dto"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
	"github.com/gofrs/uuid"
)

func (uc *UserControllers) LoginUser(w http.ResponseWriter, r *http.Request) {
	var userDTO authDTO.SignInDTO

	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		httpErr := rest_err.NewBadRequestError("Unable to parse request body")

		res, err := json.Marshal(httpErr)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(res)
		return
	}

	ctx := context.Background()

	user, err := uc.AuthServices.LoginUser(ctx, userDTO)
	if err != nil {
		if httpErr, ok := err.(*rest_err.RestErr); ok {
			res, err := json.Marshal(httpErr)
			if err != nil {
				log.Fatal(err)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			w.Write(res)
			return
		}

		return
	}

	userResponse := authDTO.AuthResponseDTO{
		ID:    uuid.UUID(user.ID),
		Email: user.Email,
		Name:  user.Name,
	}

	res, err := json.Marshal(userResponse)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(res)

	return
}
