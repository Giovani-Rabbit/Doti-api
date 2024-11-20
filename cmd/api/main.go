package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	userDTO "github.com/Giovani-Coelho/Doti-API/src/application/user/dtos"
	userServices "github.com/Giovani-Coelho/Doti-API/src/application/user/services/createUser"
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
			http.Error(w, "Unable to parse request body", http.StatusBadRequest)
			return
		}

		context := context.Background()

		addUserUseCase.CreateUser(context, user)
	})

	fmt.Println("Server is running...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
