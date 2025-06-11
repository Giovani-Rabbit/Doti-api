package container

import (
	"database/sql"

	userController "github.com/Giovani-Coelho/Doti-API/src/infra/http/controller/user"
)

type Container struct {
	DB *sql.DB
}

type IContainer interface {
	NewUserContainer(db *sql.DB) userController.IUserController
}

func NewContainer(db *sql.DB) *Container {
	return &Container{DB: db}
}
