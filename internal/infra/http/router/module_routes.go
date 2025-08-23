package router

import (
	"net/http"

	dihdl "github.com/Giovani-Coelho/Doti-API/internal/infra/di/handler"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/http/middleware"
)

func ModuleRoutes(mux *http.ServeMux, module *dihdl.ModuleHandlers) (m *http.ServeMux) {

	mux.Handle("POST /modules",
		middleware.EnsureAuth(http.HandlerFunc(
			module.Create.Execute,
		)),
	)

	mux.Handle("GET /modules",
		middleware.EnsureAuth(http.HandlerFunc(
			module.Get.Execute,
		)),
	)

	mux.Handle("PATCH /modules/{id}/rename",
		middleware.EnsureAuth(http.HandlerFunc(
			module.Rename.Execute,
		)),
	)

	mux.Handle("PATCH /modules/{id}/icon",
		middleware.EnsureAuth(http.HandlerFunc(
			module.UpdateIcon.Execute,
		)),
	)

	mux.Handle("DELETE /modules/{id}",
		middleware.EnsureAuth(http.HandlerFunc(
			module.Delete.Execute,
		)),
	)

	return
}
