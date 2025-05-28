package userUseCase

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/db/sqlc"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/repository"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
)

type FindUserByEmailUseCase struct {
	UserRepository repository.IUserRepository
}

type IFindUserByEmailUseCase interface {
	Execute(ctx context.Context, email string) (sqlc.User, error)
}

func NewFindUserByEmailUseCase(
	userRespository repository.IUserRepository,
) IFindUserByEmailUseCase {
	return &FindUserByEmailUseCase{
		UserRepository: userRespository,
	}
}

func (us *FindUserByEmailUseCase) Execute(
	ctx context.Context, email string,
) (sqlc.User, error) {
	user, err := us.UserRepository.FindUserByEmail(ctx, email)

	if err != nil {
		return sqlc.User{}, rest_err.NewNotFoundError(
			"No user with this email was found",
		)
	}

	return user, nil
}
