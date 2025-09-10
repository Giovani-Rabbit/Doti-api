package repository

import (
	"context"
	"database/sql"
	"time"

	userDomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/user"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/db/sqlc"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/mapper"
	"github.com/google/uuid"
)

type userRepository struct {
	db      *sql.DB
	queries *sqlc.Queries
}

type UserRepository interface {
	Create(ctx context.Context, user userDomain.User) (userDomain.User, error)
	CheckUserExists(ctx context.Context, email string) (bool, error)
	FindUserByEmail(ctx context.Context, email string) (userDomain.User, error)
	FindUserByEmailAndPassword(ctx context.Context, args userDomain.User) (userDomain.User, error)
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db:      db,
		queries: sqlc.New(db),
	}
}

func (ur *userRepository) Create(
	ctx context.Context,
	domainUser userDomain.User,
) (userDomain.User, error) {
	userEntity, err := ur.queries.CreateUser(ctx,
		sqlc.CreateUserParams{
			ID:        uuid.New(),
			Name:      domainUser.GetName(),
			Email:     domainUser.GetEmail(),
			Password:  domainUser.GetPassword(),
			IsAdmin:   false,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
	)

	if err != nil {
		return nil, err
	}

	return mapper.FromCreateUserRow(&userEntity), nil
}

func (ur *userRepository) CheckUserExists(
	ctx context.Context,
	email string,
) (bool, error) {
	exists, err := ur.queries.CheckUserExists(ctx, email)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func (ur *userRepository) FindUserByEmail(
	ctx context.Context,
	email string,
) (userDomain.User, error) {
	user, err := ur.queries.FindUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return mapper.FromUser(&user), nil
}

func (ur *userRepository) FindUserByEmailAndPassword(
	ctx context.Context,
	args userDomain.User,
) (userDomain.User, error) {
	user, err := ur.queries.FindUserByEmailAndPassword(ctx,
		sqlc.FindUserByEmailAndPasswordParams{
			Email:    args.GetEmail(),
			Password: args.GetPassword(),
		},
	)

	if err != nil {
		return nil, err
	}

	return mapper.FromUser(&user), nil
}
