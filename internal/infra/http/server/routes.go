package server

import (
	"database/sql"
	"net/http"

	"github.com/Giovani-Coelho/Doti-API/internal/infra/di"
)

func Routes(DB *sql.DB) (mux *http.ServeMux) {
	mux = http.NewServeMux()

	handlers := di.New(DB)

	UserRoutes(mux, handlers.User)
	ModuleRoutes(mux, handlers.Module)

	return
}
