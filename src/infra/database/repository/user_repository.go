package repository

import (
	"context"
	"database/sql"

	"github.com/Giovani-Coelho/Doti-API/src/infra/database/db/sqlc"
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
	Create(ctx context.Context) error
}

func (ur *UserRepository) Create(ctx context.Context) error {
	userEntity := sqlc.CreateUserParams{
		ID:       uuid.New(),
		Email:    "slq",
		Name:     "dawda",
		Password: "dasd123",
	}

	err := ur.Queries.CreateUser(ctx, userEntity)

	if err != nil {
		return err
	}

	return nil
}
