package router

import (
	"net/http"

	dihdl "github.com/Giovani-Coelho/Doti-API/internal/infra/di/handler"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/http/middleware"
)

func TaskRoutes(mux *http.ServeMux, task *dihdl.TaskHandlers) (m *http.ServeMux) {

	mux.Handle("POST /task",
		middleware.EnsureAuth(http.HandlerFunc(
			task.Create.Execute,
		)),
	)

	return
}
