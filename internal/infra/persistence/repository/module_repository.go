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
	DB      *sql.DB
	Queries *sqlc.Queries
}

type ModuleRepository interface {
	Create(ctx context.Context, module moduledomain.Module) (moduledomain.Module, error)
	ListModulesByUserID(ctx context.Context, userId string) ([]moduledomain.Module, error)
	UpdateModuleName(ctx context.Context, id string, name string) error
}

func NewModuleRepository(dtb *sql.DB) ModuleRepository {
	return &moduleRepository{
		DB:      dtb,
		Queries: sqlc.New(dtb),
	}
}

func (mr *moduleRepository) Create(
	ctx context.Context,
	module moduledomain.Module,
) (moduledomain.Module, error) {
	userID, err := uuid.Parse(module.GetUserId())

	if err != nil {
		return nil, err
	}

	moduleEntity, err := mr.Queries.CreateModule(ctx,
		sqlc.CreateModuleParams{
			ID:        uuid.New(),
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

func (mr *moduleRepository) ListModulesByUserID(
	ctx context.Context,
	userId string,
) ([]moduledomain.Module, error) {
	uuidUserId, err := uuid.Parse(userId)

	if err != nil {
		return nil, err
	}

	moduleEntities, err := mr.Queries.ListModuleByUserID(ctx, uuidUserId)

	if err != nil {
		return nil, err
	}

	return mapper.ConvertListModuleByUserIdRowToModules(&moduleEntities), nil
}

func (mr *moduleRepository) UpdateModuleName(
	ctx context.Context,
	id string,
	name string,
) error {
	moduleId, err := uuid.Parse(id)

	if err != nil {
		return err
	}

	err = mr.Queries.UpdateModuleName(ctx, sqlc.UpdateModuleNameParams{
		ID:   moduleId,
		Name: name,
	})

	if err != nil {
		return err
	}

	return nil
}
