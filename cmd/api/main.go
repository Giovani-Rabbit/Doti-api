package main

import (
	"fmt"
	"net/http"

	"github.com/Giovani-Coelho/Doti-API/config"
	"github.com/Giovani-Coelho/Doti-API/src/infra/http/server"
	database "github.com/Giovani-Coelho/Doti-API/src/infra/persistence"
)

var PORT = config.Env.PORT

func main() {
	conn, err := database.NewPostgresDB()

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	router := server.Routes(conn)

	fmt.Printf("Server is running on port :%d", PORT)

	http.ListenAndServe(
		fmt.Sprintf(":%d", PORT),
		router,
	)
}
