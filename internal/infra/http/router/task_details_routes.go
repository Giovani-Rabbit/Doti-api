package router

import (
	"net/http"

	dihdl "github.com/Giovani-Coelho/Doti-API/internal/infra/di/handler"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/http/middleware"
)

func TaskDetailsRoutes(mux *http.ServeMux, details *dihdl.TaskDetailsHandler) (m *http.ServeMux) {

	mux.Handle("PATCH /tasks/{id}/description",
		middleware.EnsureAuth(http.HandlerFunc(
			details.UpdateDescription.Execute,
		)),
	)

	mux.Handle("PATCH /tasks/{id}/pomodoro-target",
		middleware.EnsureAuth(http.HandlerFunc(
			details.UpdatePomodoroTarget.Execute,
		)),
	)

	return
}
