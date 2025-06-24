package authServices

import (
	"context"

	userdto "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/user/dtos"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/db/sqlc"
	"github.com/Giovani-Coelho/Doti-API/src/pkg/auth"
	"github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/encrypt"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
)

func (as *AuthServices) LoginUser(
	ctx context.Context,
	userDTO userdto.SignInDTO,
) (sqlc.User, string, error) {
	encryptPassword := sqlc.FindUserByEmailAndPasswordParams{
		Email:    userDTO.Email,
		Password: encrypt.EncryptPassword(userDTO.Password),
	}

	user, err := as.UserRepository.FindUserByEmailAndPassword(ctx, encryptPassword)

	if err != nil {
		return sqlc.User{}, "", rest_err.NewBadRequestError(
			"UNKNOWN",
			"Invalid login credentials",
		)
	}

	token, err := auth.GenerateToken(user)

	if err != nil {
		return sqlc.User{}, "", rest_err.NewBadRequestError(
			"UNKNOWN",
			"Error on generate Token",
		)
	}

	return user, token, nil
}
