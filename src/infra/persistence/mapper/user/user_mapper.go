package userMapper

import (
	userDomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/db/sqlc"
)

func FromCreateUserRow(
	user *sqlc.CreateUserRow,
) (domain userDomain.IUserDomain) {
	return userDomain.NewUserDomain(
		user.ID.String(),
		user.Name,
		user.Email,
		"",
		user.CreatedAt,
		user.UpdatedAt,
	)
}
