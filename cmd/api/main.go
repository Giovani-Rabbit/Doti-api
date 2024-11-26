package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Giovani-Coelho/Doti-API/src/infra/database"
	"github.com/Giovani-Coelho/Doti-API/src/infra/http/server"
)

func main() {
	conn, err := database.NewPostgresDB()

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	router := server.Routes(conn)

	fmt.Println("Server is running...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
