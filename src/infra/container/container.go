package container

import (
	"database/sql"

	userhandler "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/user"
)

type Container struct {
	DB *sql.DB
}

type IContainer interface {
	NewUserContainer(db *sql.DB) userhandler.IUserHandler
}

func NewContainer(db *sql.DB) *Container {
	return &Container{DB: db}
}
