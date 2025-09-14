package taskcase

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	taskdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	"go.uber.org/zap"
)

type GetByModule interface {
	Execute(ctx context.Context, moduleId int32) ([]taskdomain.Task, error)
}

type getByModule struct {
	taskRepo   repository.TaskRepository
	moduleRepo repository.ModuleRepository
}

func NewGetTasksByModuleId(
	tr repository.TaskRepository,
	mr repository.ModuleRepository,
) GetByModule {
	return &getByModule{
		taskRepo:   tr,
		moduleRepo: mr,
	}
}

func (gt *getByModule) Execute(
	ctx context.Context, moduleId int32,
) ([]taskdomain.Task, error) {
	logger.Info("Init get tasks by module id",
		zap.String("journey", "getTaskByModuleId"),
		zap.Int32("moduleId", moduleId))

	if moduleId <= 0 {
		logger.Error("Invalid module id", nil,
			zap.String("journey", "getTaskByModuleId"),
			zap.Int32("moduleId", moduleId))

		return nil, moduledomain.ErrInvalidModuleID()
	}

	moduleExists, err := gt.moduleRepo.CheckExistsById(ctx, moduleId)
	if err != nil {
		logger.Error("Failed to check if module exists", err,
			zap.String("journey", "getTaskByModuleId"),
			zap.Int32("moduleId", moduleId))

		return nil, moduledomain.ErrCheckingIfModuleExists()
	}

	if !moduleExists {
		logger.Error("Module not founded", nil,
			zap.String("journey", "getTaskByModuleId"),
			zap.Int32("moduleId", moduleId))

		return nil, moduledomain.ErrCouldNotFindModuleByID()
	}

	tasks, err := gt.taskRepo.ListByModuleId(ctx, moduleId)
	if err != nil {
		logger.Error("Failed to get task list", err,
			zap.String("journey", "getTaskByModuleId"),
			zap.Int32("moduleId", moduleId))

		return nil, taskdomain.ErrCouldNotListTasks()
	}

	logger.Info("List of tasks obtained",
		zap.String("journey", "getTaskByModuleId"),
		zap.Int32("moduleId", moduleId))

	return tasks, nil
}
