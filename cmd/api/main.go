package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Giovani-Coelho/Doti-API/src/infra/database"
	db "github.com/Giovani-Coelho/Doti-API/src/infra/database/db/sqlc"
	user_repository "github.com/Giovani-Coelho/Doti-API/src/infra/database/repository/user"
	"github.com/gorilla/mux"
)

func main() {
	conn, err := database.NewPostgresDB()

	if err != nil {
		panic(err)
	}

	defer conn.Close()
	queries := db.New(conn)

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		userRepository := user_repository.NewUserRepository(queries)

		users, err := userRepository.GetUsers(r.Context())

		if err != nil {
			json.NewEncoder(w).Encode("Error na busca de users")
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})

	fmt.Println("Server is running...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
