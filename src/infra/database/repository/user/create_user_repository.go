package user_repository

import (
	"context"

	entity "github.com/Giovani-Coelho/Doti-API/src/infra/database/entity/user"
	"github.com/gofrs/uuid"
)

func (ur *UserRepository) GetUsers(ctx context.Context) ([]entity.User, error) {
	users, err := ur.queries.GetUsers(ctx)

	if err != nil {
		return nil, err
	}

	var usersEntity []entity.User

	for _, user := range users {
		userEntity := entity.User{
			ID:       uuid.UUID(user.ID),
			Email:    user.Email,
			Name:     user.Name,
			Password: user.Password,
		}

		usersEntity = append(usersEntity, userEntity)
	}

	return usersEntity, nil
}
