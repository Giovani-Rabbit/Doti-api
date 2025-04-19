package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Giovani-Coelho/Doti-API/src/infra/database"
	"github.com/Giovani-Coelho/Doti-API/src/infra/http/server"
	"github.com/rs/cors"
)

func main() {
	conn, err := database.NewPostgresDB()

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	router := server.Routes(conn)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	fmt.Println("Server is running...")

	log.Fatal(http.ListenAndServe(":8080", handler))
}
