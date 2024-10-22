package user_repository

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/src/infra/database/db/sqlc"
	entity "github.com/Giovani-Coelho/Doti-API/src/infra/database/entity/user"
)

func NewUserRepository(q *sqlc.Queries) IUserRepository {
	return &UserRepository{q}
}

type UserRepository struct {
	queries *sqlc.Queries
}

type IUserRepository interface {
	GetUsers(ctx context.Context) ([]entity.User, error)
}
