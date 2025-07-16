package repository

import (
	"context"
	"database/sql"
	"time"

	userDomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/db/sqlc"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/mapper"
	"github.com/google/uuid"
)

func NewUserRepository(dtb *sql.DB) IUserRepository {
	return &UserRepository{
		DB:      dtb,
		Queries: sqlc.New(dtb),
	}
}

type UserRepository struct {
	DB      *sql.DB
	Queries *sqlc.Queries
}

type IUserRepository interface {
	Create(ctx context.Context, user userDomain.IUserDomain) (userDomain.IUserDomain, error)
	CheckUserExists(ctx context.Context, email string) (bool, error)
	FindUserByEmail(ctx context.Context, email string) (userDomain.IUserDomain, error)
	FindUserByEmailAndPassword(ctx context.Context, args userDomain.IUserDomain) (userDomain.IUserDomain, error)
}

func (ur *UserRepository) Create(
	ctx context.Context,
	domainUser userDomain.IUserDomain,
) (userDomain.IUserDomain, error) {
	userEntity, err := ur.Queries.CreateUser(ctx,
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

func (ur *UserRepository) CheckUserExists(
	ctx context.Context,
	email string,
) (bool, error) {
	exists, err := ur.Queries.CheckUserExists(ctx, email)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func (ur *UserRepository) FindUserByEmail(
	ctx context.Context,
	email string,
) (userDomain.IUserDomain, error) {
	user, err := ur.Queries.FindUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return mapper.FromUser(&user), nil
}

func (ur *UserRepository) FindUserByEmailAndPassword(
	ctx context.Context,
	args userDomain.IUserDomain,
) (userDomain.IUserDomain, error) {
	user, err := ur.Queries.FindUserByEmailAndPassword(ctx,
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
