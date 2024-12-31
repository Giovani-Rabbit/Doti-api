package userController

import (
	"net/http"

	userServices "github.com/Giovani-Coelho/Doti-API/src/application/user/services"
)

type UserControllers struct {
	UserServices userServices.IUserServices
}

type IUserControllers interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

func NewUserControllers(
	userServices userServices.IUserServices,
) IUserControllers {
	return &UserControllers{
		UserServices: userServices,
	}
}
