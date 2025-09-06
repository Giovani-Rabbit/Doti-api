package modulecase

import (
	"context"
	"strconv"
	"strings"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	"go.uber.org/zap"
)

type updateIcon struct {
	moduleRepository repository.ModuleRepository
}

type UpdateIcon interface {
	Execute(ctx context.Context, moduleId string, icon string) error
}

func NewUpdateModuleIconUseCase(
	moduleRepo repository.ModuleRepository,
) UpdateIcon {
	return &updateIcon{
		moduleRepository: moduleRepo,
	}
}

func (ui *updateIcon) Execute(
	ctx context.Context, moduleId string, icon string,
) error {
	logger.Info("Init change module icon",
		zap.String("journey", "changeIcon"),
		zap.String("newIcon", icon),
		zap.String("moduleId", moduleId))

	if strings.TrimSpace(moduleId) == "" {
		logger.Error("module id is empty", nil,
			zap.String("journey", "changeIcon"),
			zap.String("moduleId", moduleId))

		return moduledomain.ErrInvalidModuleID()
	}

	intId, err := strconv.ParseInt(moduleId, 10, 32)
	if err != nil {
		logger.Error("Id is not a int", err,
			zap.String("journey", "changeIcon"),
			zap.String("moduleId", moduleId))

		return moduledomain.ErrInvalidModuleID()
	}

	int32ModuleId := int32(intId)

	if strings.TrimSpace(icon) == "" {
		logger.Error("The new icon is empty", nil,
			zap.String("journey", "changeIcon"),
			zap.String("newIcon", icon))

		return moduledomain.ErrNewModuleIconIsEmpty()
	}

	icon = strings.TrimSpace(icon)

	err = ui.moduleRepository.UpdateIcon(ctx, int32ModuleId, icon)
	if err != nil {
		logger.Error("Failed to update module icon", err,
			zap.String("journey", "changeIcon"),
			zap.Int32("moduleId", int32ModuleId),
			zap.String("newIcon", icon))

		return moduledomain.ErrCouldNotPersistModule(err)
	}

	logger.Info("Module icon updated",
		zap.String("journey", "changeIcon"),
		zap.String("newIcon", icon),
		zap.Int32("moduleId", int32ModuleId))

	return nil
}
