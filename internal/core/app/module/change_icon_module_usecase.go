package modulecase

import (
	"context"
	"strings"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	val "github.com/Giovani-Coelho/Doti-API/internal/pkg/validator"
	"go.uber.org/zap"
)

type changeIcon struct {
	moduleRepository repository.ModuleRepository
}

type ChangeIcon interface {
	Execute(ctx context.Context, moduleId string, icon string) error
}

func NewChangeModuleIconUseCase(
	moduleRepo repository.ModuleRepository,
) ChangeIcon {
	return &changeIcon{
		moduleRepository: moduleRepo,
	}
}

func (ci *changeIcon) Execute(
	ctx context.Context, moduleId string, icon string,
) error {
	logger.Info("Init change module icon",
		zap.String("journey", "changeIcon"),
	)

	if !val.IsValidUUID(moduleId) {
		logger.Error("Invalid module id", nil,
			zap.String("journey", "changeIcon"))

		return moduledomain.ErrInvalidModuleID()
	}

	if strings.TrimSpace(icon) == "" {
		logger.Error("Error, the new icon is empty", nil,
			zap.String("journey", "changeIcon"),
		)

		return moduledomain.ErrNewModuleIconIsEmpty()
	}

	icon = strings.TrimSpace(icon)

	err := ci.moduleRepository.UpdateIcon(ctx, moduleId, icon)

	if err != nil {
		logger.Error("Error on update module icon", err,
			zap.String("journey", "changeIcon"),
		)

		return moduledomain.ErrCouldNotPersistModule(err)
	}

	logger.Info("Icon updated successfully",
		zap.String("journey", "changeIcon"),
	)

	return nil
}
