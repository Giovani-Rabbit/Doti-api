package userServices

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/src/infra/database/db/sqlc"
	"github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/encrypt"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
)

func (us *UserServices) findUserByEmailAndPassword(
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
