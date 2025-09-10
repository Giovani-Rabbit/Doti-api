package repository

import (
	"context"
	"database/sql"
	"time"

	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/db/sqlc"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/mapper"
	"github.com/google/uuid"
)

type moduleRepository struct {
	db      *sql.DB
	queries *sqlc.Queries
}

type ModuleRepository interface {
	CheckExistsById(ctx context.Context, id int32) (bool, error)
	Create(ctx context.Context, module moduledomain.Module) (moduledomain.Module, error)
	DeleteModule(ctx context.Context, id int32) error
	ListModulesByUserID(ctx context.Context, userId string) ([]moduledomain.Module, error)
	UpdateModuleName(ctx context.Context, id int32, name string) error
	UpdateIcon(ctx context.Context, id int32, icon string) error
}

func NewModuleRepository(dtb *sql.DB) ModuleRepository {
	return &moduleRepository{
		db:      dtb,
		queries: sqlc.New(dtb),
	}
}

func (mr *moduleRepository) CheckExistsById(
	ctx context.Context, id int32,
) (bool, error) {
	exists, err := mr.queries.CheckModuleExists(ctx, id)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func (mr *moduleRepository) Create(
	ctx context.Context,
	module moduledomain.Module,
) (moduledomain.Module, error) {
	userID, err := uuid.Parse(module.GetUserId())

	if err != nil {
		return nil, err
	}

	moduleEntity, err := mr.queries.CreateModule(ctx,
		sqlc.CreateModuleParams{
			UserID:    userID,
			Name:      module.GetName(),
			IsOpen:    false,
			Icon:      module.GetIcon(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	)

	if err != nil {
		return nil, err
	}

	return mapper.ConvertCreateModuleRowToModule(&moduleEntity), nil
}

func (mr *moduleRepository) DeleteModule(
	ctx context.Context, id int32,
) error {
	err := mr.queries.DeleteModule(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

func (mr *moduleRepository) ListModulesByUserID(
	ctx context.Context,
	userId string,
) ([]moduledomain.Module, error) {
	uuidUserId, err := uuid.Parse(userId)

	if err != nil {
		return nil, err
	}

	moduleEntities, err := mr.queries.ListModuleByUserID(ctx, uuidUserId)

	if err != nil {
		return nil, err
	}

	return mapper.ConvertListModuleByUserIdRowToModules(&moduleEntities), nil
}

func (mr *moduleRepository) UpdateModuleName(
	ctx context.Context,
	id int32,
	name string,
) error {
	err := mr.queries.UpdateModuleName(ctx, sqlc.UpdateModuleNameParams{
		ID:   id,
		Name: name,
	})

	if err != nil {
		return err
	}

	return nil
}

func (mr *moduleRepository) UpdateIcon(
	ctx context.Context, id int32, icon string,
) error {
	err := mr.queries.UpdateIcon(ctx, sqlc.UpdateIconParams{
		ID:   id,
		Icon: icon,
	})

	if err != nil {
		return err
	}

	return nil
}
