package container

import (
	"database/sql"

	authhandler "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/auth"
	modulehandler "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/module"
	userhandler "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/user"
)

type container struct {
	DB *sql.DB
}

type Container interface {
	NewUser() userhandler.UserHandler
	NewAuth() authhandler.AuthHandler
	NewModule() modulehandler.ModuleHandler
}

func NewContainer(db *sql.DB) Container {
	return &container{DB: db}
}
