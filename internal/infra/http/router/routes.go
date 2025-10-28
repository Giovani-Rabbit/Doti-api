package router

import (
	"database/sql"
	"net/http"

	"github.com/Giovani-Coelho/Doti-API/internal/infra/di"
)

func Routes(DB *sql.DB) (mux *http.ServeMux) {
	mux = http.NewServeMux()

	handlers := di.New(DB)

	ModuleRoutes(mux, handlers.Module)
	TaskRoutes(mux, handlers.Task)
	TaskDetailsRoutes(mux, handlers.TaskDetails)
	UserRoutes(mux, handlers.User)

	return
}
