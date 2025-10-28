package router

import (
	"net/http"

	dihdl "github.com/Giovani-Coelho/Doti-API/internal/infra/di/handler"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/http/middleware"
)

func TaskDetailsRoutes(mux *http.ServeMux, details *dihdl.TaskDetailsHandler) (m *http.ServeMux) {

	mux.Handle("PATCH /tasks/{id}/details",
		middleware.EnsureAuth(http.HandlerFunc(
			details.UpdateDescription.Execute,
		)),
	)

	return
}
