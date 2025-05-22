package server

import (
	"database/sql"
	"net/http"
)

func Routes(DB *sql.DB) *http.ServeMux {
	router := http.NewServeMux()

	return router
}
