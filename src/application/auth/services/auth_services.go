package authServices

import (
	"github.com/Giovani-Coelho/Doti-API/src/infra/database/repository"
)

type AuthServices struct {
}

type IAuthServices interface {
}

func NewAuthServices(
	userRepository repository.IUserRepository,
) {

}
