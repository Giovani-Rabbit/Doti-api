package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	userDTO "github.com/Giovani-Coelho/Doti-API/src/application/user/dtos"
	userServices "github.com/Giovani-Coelho/Doti-API/src/application/user/services/createUser"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/handlers/http"
	"github.com/Giovani-Coelho/Doti-API/src/infra/database"
	"github.com/Giovani-Coelho/Doti-API/src/infra/database/repository"
	"github.com/gorilla/mux"
)

func main() {
	conn, err := database.NewPostgresDB()

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	useRepository := repository.NewUserRepository(conn)
	addUserUseCase := userServices.NewCreateUserService(useRepository)

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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

		err = addUserUseCase.CreateUser(ctx, user)
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
	})

	fmt.Println("Server is running...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func WriteJSON(w http.ResponseWriter, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)
	w.Write(response)
}
