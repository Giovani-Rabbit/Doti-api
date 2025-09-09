package modulecase

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	val "github.com/Giovani-Coelho/Doti-API/internal/pkg/validator"
	"go.uber.org/zap"
)

type GetByUser interface {
	Execute(ctx context.Context, userId string) (
		[]moduledomain.Module, error)
}

type getByUser struct {
	moduleRepo repository.ModuleRepository
}

func NewGetModulesUseCase(
	moduleRepo repository.ModuleRepository,
) GetByUser {
	return &getByUser{
		moduleRepo: moduleRepo,
	}
}

func (gm *getByUser) Execute(
	ctx context.Context, userId string,
) ([]moduledomain.Module, error) {
	logger.Info("Init get modules",
		zap.String("journey", "getModule"),
		zap.String("userId", userId))

	if !val.IsValidUUID(userId) {
		logger.Error("User id is not a uuid", nil,
			zap.String("userId", userId))

		return nil, resp.NewInvalidUUID()
	}

	modules, err := gm.moduleRepo.ListModulesByUserID(ctx, userId)
	if err != nil {
		logger.Error("Error in repository", err,
			zap.String("journey", "getModule"),
			zap.String("userId", userId))

		return nil, moduledomain.ErrCouldNotPersistModule(err)
	}

	logger.Info("List of modules obtained",
		zap.String("journey", "getModule"),
		zap.Int("ModulesLength", len(modules)),
		zap.String("userId", userId))

	return modules, nil
}
