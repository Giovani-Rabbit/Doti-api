package container

import (
	"database/sql"

	authhandler "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/auth"
	userhandler "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/user"
)

type Container struct {
	DB *sql.DB
}

type IContainer interface {
	NewUserContainer() userhandler.IUserHandler
	NewAuthContainer() authhandler.IAuthHandler
}

func NewContainer(db *sql.DB) *Container {
	return &Container{DB: db}
}
