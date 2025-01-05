package userServices

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/src/infra/database/db/sqlc"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
)

func (us *UserServices) FindUserByEmail(
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
