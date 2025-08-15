package router

import (
	"net/http"

	dihdl "github.com/Giovani-Coelho/Doti-API/internal/infra/di/handler"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/http/middleware"
)

func ModuleRoutes(mux *http.ServeMux, module *dihdl.ModuleHandlers) (m *http.ServeMux) {

	mux.Handle("POST /module",
		middleware.EnsureAuth(http.HandlerFunc(
			module.Create.Execute,
		)),
	)

	mux.Handle("GET /module",
		middleware.EnsureAuth(http.HandlerFunc(
			module.Get.Execute,
		)),
	)

	mux.Handle("PATCH /module/{id}",
		middleware.EnsureAuth(http.HandlerFunc(
			module.Rename.Execute,
		)),
	)

	mux.Handle("DELETE /module/{id}",
		middleware.EnsureAuth(http.HandlerFunc(
			module.Delete.Execute,
		)),
	)

	return
}
