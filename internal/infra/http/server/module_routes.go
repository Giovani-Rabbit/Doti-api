package server

import (
	"net/http"

	"github.com/Giovani-Coelho/Doti-API/internal/infra/container"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/http/middleware"
)

func ModuleRoutes(mux *http.ServeMux, module *container.ModuleHandler) (m *http.ServeMux) {

	mux.Handle("POST /module",
		middleware.EnsureAuth(http.HandlerFunc(
			module.Create.Handler,
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
