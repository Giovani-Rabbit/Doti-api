package authServices

import (
	"context"

	authDTO "github.com/Giovani-Coelho/Doti-API/src/application/auth/dto"
	"github.com/Giovani-Coelho/Doti-API/src/infra/database/db/sqlc"
	"github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/encrypt"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
)

func (as *AuthServices) LoginUser(
	ctx context.Context,
	userDTO authDTO.SignInDTO,
) (sqlc.User, error) {
	encryptPassword := sqlc.FindUserByEmailAndPasswordParams{
		Email:    userDTO.Email,
		Password: encrypt.EncryptPassword(userDTO.Password),
	}

	user, err := as.UserRepository.FindUserByEmailAndPassword(ctx, encryptPassword)

	if err != nil {
		return sqlc.User{}, rest_err.NewBadRequestError(
			"Invalid login credentials",
		)
	}

	return user, nil
}
