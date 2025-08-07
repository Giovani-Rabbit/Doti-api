package modulecase

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	apperr "github.com/Giovani-Coelho/Doti-API/internal/core/app/errors"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	"go.uber.org/zap"
)

type renameModuleUsecase struct {
	ModuleRepo repository.ModuleRepository
}

type RenameModuleUseCase interface {
	Execute(ctx context.Context, id string, name string) error
}

func NewRenameModuleUseCase(
	moduleRepo repository.ModuleRepository,
) RenameModuleUseCase {
	return &renameModuleUsecase{
		ModuleRepo: moduleRepo,
	}
}

func (rm *renameModuleUsecase) Execute(
	ctx context.Context,
	id string, name string,
) error {
	logger.Info("Init rename module",
		zap.String("journey", "renameModule"),
	)

	err := rm.ModuleRepo.UpdateModuleName(ctx, id, name)

	if err != nil {
		logger.Error("Error on change module name", err,
			zap.String("journey", "renameModule"),
		)

		return apperr.ErrRenamingModule(err)
	}

	logger.Info("Module renamed successfully",
		zap.String("journey", "renameModule"),
	)

	return nil
}
