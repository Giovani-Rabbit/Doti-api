package main

import (
	"fmt"
	"log"
	"net/http"

	user_controller "github.com/Giovani-Coelho/Doti-API/src/infra/http/controller/user"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", user_controller.CreateUser)

	fmt.Println("Server running...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
