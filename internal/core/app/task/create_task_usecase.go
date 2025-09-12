package taskcase

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	taskdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	"go.uber.org/zap"
)

type Create interface {
	Execute(ctx context.Context, task taskdomain.Task) (taskdomain.Task, error)
}

type create struct {
	taskRepo   repository.TaskRepository
	moduleRepo repository.ModuleRepository
}

func NewCreateTaskUseCase(
	tr repository.TaskRepository,
	mr repository.ModuleRepository,
) Create {
	return &create{
		taskRepo:   tr,
		moduleRepo: mr,
	}
}

func (ct *create) Execute(
	ctx context.Context,
	task taskdomain.Task,
) (taskdomain.Task, error) {
	logger.Info("Init create task",
		zap.String("journey", "createTask"),
		zap.String("taskName", task.Name()),
		zap.Int32("moduleId", task.ModuleId()))

	if !task.IsValidToCreate() {
		logger.Error("Invalid Task Fields", nil,
			zap.String("journey", "createTask"),
			zap.String("taskName", task.Name()),
			zap.Int32("moduleId", task.ModuleId()),
			zap.Int32("taskPosition", task.Position()))

		return nil, taskdomain.ErrInvalidFields()
	}

	moduleExists, err := ct.moduleRepo.CheckExistsById(ctx, task.ModuleId())
	if err != nil {
		logger.Error("Failed to check if module exists", err,
			zap.String("journey", "createTask"),
			zap.Int32("moduleId", task.ModuleId()))

		return nil, moduledomain.ErrCheckingIfModuleExists()
	}

	if !moduleExists {
		logger.Error("Module not founded", nil,
			zap.String("journey", "createTask"),
			zap.Int32("moduleId", task.ModuleId()))

		return nil, moduledomain.ErrCouldNotFindModuleByID()
	}

	createdTask, err := ct.taskRepo.Create(ctx, task)
	if err != nil {
		logger.Error("Failed to persist task", err,
			zap.String("journey", "createTask"),
			zap.String("taskName", task.Name()),
			zap.Int32("moduleId", task.ModuleId()),
			zap.Int32("position", task.Position()))

		return nil, taskdomain.ErrCouldNotToCreate()
	}

	logger.Info("Task created successfully",
		zap.String("journey", "createTask"))

	return createdTask, nil
}
