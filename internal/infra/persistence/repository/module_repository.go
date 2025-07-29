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

type ModuleRepository struct {
	DB      *sql.DB
	Queries *sqlc.Queries
}

type IModuleRepository interface {
	Create(ctx context.Context, module moduledomain.IModuleDomain) (moduledomain.IModuleDomain, error)
}

func NewModuleRepository(dtb *sql.DB) IModuleRepository {
	return &ModuleRepository{
		DB:      dtb,
		Queries: sqlc.New(dtb),
	}
}

func (mr *ModuleRepository) Create(
	ctx context.Context,
	module moduledomain.IModuleDomain,
) (moduledomain.IModuleDomain, error) {
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

	return mapper.FromCreateModuleRow(&moduleEntity), nil
}
