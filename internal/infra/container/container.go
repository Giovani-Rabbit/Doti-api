package container

import (
	"database/sql"

	authhandler "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/auth"
	modulehandler "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/module"
	userhandler "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/user"
)

type Container struct {
	DB *sql.DB
}

type IContainer interface {
	NewUserContainer() userhandler.IUserHandler
	NewAuthContainer() authhandler.IAuthHandler
	NewModuleContainer() modulehandler.IModuleHandler
}

func NewContainer(db *sql.DB) *Container {
	return &Container{DB: db}
}
