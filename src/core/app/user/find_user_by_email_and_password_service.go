package userServices

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/db/sqlc"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/repository"
	"github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/encrypt"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
)

type FindUserByEmailAndPasswordService struct {
	UserRepository repository.IUserRepository
}

type IFindUserByEmailAndPasswordService interface {
	Execute(ctx context.Context, email string, password string) (sqlc.User, error)
}

func NewFindUserByEmailAndPasswordService(
	userRepository repository.IUserRepository,
) IFindUserByEmailAndPasswordService {
	return &FindUserByEmailAndPasswordService{
		UserRepository: userRepository,
	}
}

func (us *FindUserByEmailAndPasswordService) Execute(
	ctx context.Context,
	email string,
	password string,
) (sqlc.User, error) {
	bindArgs := sqlc.FindUserByEmailAndPasswordParams{
		Email:    email,
		Password: encrypt.EncryptPassword(password),
	}

	user, err := us.UserRepository.FindUserByEmailAndPassword(ctx, bindArgs)

	if err != nil {
		return sqlc.User{}, rest_err.NewNotFoundError(
			"Unable to find a user with these arguments.",
		)
	}

	return user, nil
}
