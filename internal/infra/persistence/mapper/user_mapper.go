package mapper

import (
	userDomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/user"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/db/sqlc"
)

func FromCreateUserRow(
	user *sqlc.CreateUserRow,
) (domain userDomain.User) {
	return userDomain.New(
		user.ID.String(),
		user.Name,
		user.Email,
		"",
		user.CreatedAt,
		user.UpdatedAt,
	)
}

func FromUser(
	user *sqlc.User,
) (domain userDomain.User) {
	return userDomain.New(
		user.ID.String(),
		user.Name,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	)
}
