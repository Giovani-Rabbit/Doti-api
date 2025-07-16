package repository

import (
	"context"
	"database/sql"

	moduledomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/module"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/db/sqlc"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/mapper"
	"github.com/google/uuid"
)

type ModuleRepository struct {
	DB      *sql.DB
	Queries *sqlc.Queries
}

type IModuleRepository interface {
	Create(ctx context.Context, module moduledomain.IModuleDomain) (moduledomain.IModuleDomain, error)
}

func New(dtb *sql.DB) IModuleRepository {
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
			IsOpen:    module.GetIsOpen(),
			Icon:      module.GetIcon(),
			CreatedAt: module.GetCreateAt(),
			UpdatedAt: module.GetUpdatedAt(),
		},
	)

	if err != nil {
		return nil, err
	}

	return mapper.FromCreateModuleRow(&moduleEntity), nil
}
