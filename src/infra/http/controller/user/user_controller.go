package userController

import (
	"net/http"

	authServices "github.com/Giovani-Coelho/Doti-API/src/application/auth/services"
	userServices "github.com/Giovani-Coelho/Doti-API/src/application/user/services"
)

type UserControllers struct {
	UserServices userServices.IUserServices
	AuthServices authServices.IAuthServices
}

type IUserControllers interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
}

func NewUserControllers(
	userServices userServices.IUserServices,
	authServices authServices.IAuthServices,
) IUserControllers {
	return &UserControllers{
		UserServices: userServices,
		AuthServices: authServices,
	}
}
