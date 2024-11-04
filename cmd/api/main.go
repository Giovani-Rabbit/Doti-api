package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Giovani-Coelho/Doti-API/src/infra/database"
	"github.com/Giovani-Coelho/Doti-API/src/infra/database/repository"
	useCase "github.com/Giovani-Coelho/Doti-API/src/useCase/user"
	"github.com/gorilla/mux"
)

func main() {
	conn, err := database.NewPostgresDB()

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	context := context.Background()

	useRepository := repository.NewUserRepository(conn)
	addUserUseCase := useCase.NewCreateUserUseCase(useRepository)

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		addUserUseCase.Execute(context)
	})

	fmt.Println("Server is running...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
