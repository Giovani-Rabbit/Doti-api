package repository

import (
	"context"
	"database/sql"

	userDTO "github.com/Giovani-Coelho/Doti-API/src/core/user/dtos"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/db/sqlc"
	"github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/encrypt"
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
	Create(ctx context.Context, userDto userDTO.CreateUserDTO) error
	CheckUserExists(ctx context.Context, email string) (bool, error)
	FindUserByEmail(ctx context.Context, email string) (sqlc.User, error)
	FindUserByEmailAndPassword(
		ctx context.Context,
		args sqlc.FindUserByEmailAndPasswordParams,
	) (sqlc.User, error)
}

func (ur *UserRepository) Create(ctx context.Context, userDTO userDTO.CreateUserDTO) error {
	userEntity := sqlc.CreateUserParams{
		ID:       uuid.New(),
		Name:     userDTO.Name,
		Email:    userDTO.Email,
		Password: encrypt.EncryptPassword(userDTO.Password),
	}

	err := ur.Queries.CreateUser(ctx, userEntity)

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) CheckUserExists(ctx context.Context, email string) (bool, error) {
	exists, err := ur.Queries.CheckUserExists(ctx, email)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func (ur *UserRepository) FindUserByEmail(ctx context.Context, email string) (sqlc.User, error) {
	user, err := ur.Queries.FindUserByEmail(ctx, email)

	if err != nil {
		return sqlc.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) FindUserByEmailAndPassword(
	ctx context.Context,
	args sqlc.FindUserByEmailAndPasswordParams,
) (sqlc.User, error) {
	user, err := ur.Queries.FindUserByEmailAndPassword(ctx, args)

	if err != nil {
		return sqlc.User{}, err
	}

	return user, nil
}
