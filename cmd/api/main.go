package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Giovani-Coelho/Doti-API/config"
	"github.com/Giovani-Coelho/Doti-API/src/infra/http/server"
	database "github.com/Giovani-Coelho/Doti-API/src/infra/persistence"
)

func main() {
	// corsOrigin := config.Env.CorsOrigin
	serverPort := config.Env.ServerPort

	conn, err := database.NewPostgresDB()

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	router := server.Routes(conn)

	fmt.Printf("Server is running on port :%d", serverPort)

	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%d", serverPort),
			router,
		),
	)
}
