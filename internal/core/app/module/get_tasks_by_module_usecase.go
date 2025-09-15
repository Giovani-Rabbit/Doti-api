package modulecase

import (
	"context"
	"strconv"
	"strings"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	taskdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	"go.uber.org/zap"
)

type GetTasks interface {
	Execute(ctx context.Context, moduleId string) ([]taskdomain.Task, error)
}

type getTasks struct {
	taskRepo   repository.TaskRepository
	moduleRepo repository.ModuleRepository
}

func NewGetTasksByModuleId(
	tr repository.TaskRepository,
	mr repository.ModuleRepository,
) GetTasks {
	return &getTasks{
		taskRepo:   tr,
		moduleRepo: mr,
	}
}

func (gt *getTasks) Execute(
	ctx context.Context, moduleId string,
) ([]taskdomain.Task, error) {
	logger.Info("Init get tasks by module id",
		zap.String("journey", "getTaskByModuleId"),
		zap.String("moduleId", moduleId))

	if strings.TrimSpace(moduleId) == "" {
		logger.Error("module id is empty", nil,
			zap.String("journey", "getTaskByModuleId"),
			zap.String("moduleId", moduleId))

		return nil, moduledomain.ErrInvalidModuleID()
	}

	intId, err := strconv.ParseInt(moduleId, 10, 32)
	if err != nil {
		logger.Error("Invalid, id is not a int", err,
			zap.String("journey", "getTaskByModuleId"),
			zap.String("moduleId", moduleId))

		return nil, moduledomain.ErrInvalidModuleID()
	}

	int32Id := int32(intId)

	if int32Id <= 0 {
		logger.Error("Invalid module id", nil,
			zap.String("journey", "getTaskByModuleId"),
			zap.String("moduleId", moduleId))

		return nil, moduledomain.ErrInvalidModuleID()
	}

	moduleExists, err := gt.moduleRepo.CheckExistsById(ctx, int32Id)
	if err != nil {
		logger.Error("Failed to check if module exists", err,
			zap.String("journey", "getTaskByModuleId"),
			zap.Int32("moduleId", int32Id))

		return nil, moduledomain.ErrCheckingIfModuleExists()
	}

	if !moduleExists {
		logger.Error("Module not founded", nil,
			zap.String("journey", "getTaskByModuleId"),
			zap.Int32("moduleId", int32Id))

		return nil, moduledomain.ErrCouldNotFindModuleByID()
	}

	tasks, err := gt.taskRepo.ListByModuleId(ctx, int32Id)
	if err != nil {
		logger.Error("Failed to get task list", err,
			zap.String("journey", "getTaskByModuleId"),
			zap.Int32("moduleId", int32Id))

		return nil, taskdomain.ErrCouldNotListTasks()
	}

	logger.Info("List of tasks obtained",
		zap.String("journey", "getTaskByModuleId"))

	return tasks, nil
}
