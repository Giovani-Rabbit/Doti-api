package server

import (
	"database/sql"
	"net/http"

	"github.com/Giovani-Coelho/Doti-API/internal/infra/container"
)

func Routes(DB *sql.DB) (mux *http.ServeMux) {
	mux = http.NewServeMux()

	c := container.Setup(DB)

	UserRoutes(mux, c.User)
	ModuleRoutes(mux, c.Module)

	return
}
