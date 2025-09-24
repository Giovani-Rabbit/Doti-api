package taskcase

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	taskdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task"
	taskdto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/task/dtos"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	"go.uber.org/zap"
)

type UpdatePosition interface {
	Execute(ctx context.Context, params *taskdto.UpdatePositionDTO) error
}

type updatePosition struct {
	taskRepository repository.TaskRepository
}

func NewUpdateTaskPosition(
	tr repository.TaskRepository,
) UpdatePosition {
	return &updatePosition{
		taskRepository: tr,
	}
}

func (tp *updatePosition) Execute(
	ctx context.Context, params *taskdto.UpdatePositionDTO,
) error {
	logger.Info("Update task position",
		zap.String("journey", "updateTaskPosition"))

	if len(params.Tasks) != 2 {
		logger.Error("The tasks number must be equal to 2", nil,
			zap.Int("tasks_quantity", len(params.Tasks)),
			zap.String("journey", "updateTaskPosition"))

		return taskdomain.ErrInvalidTaskQuantity()
	}

	if params.Tasks[0].Position == params.Tasks[1].Position {
		logger.Error("The task position cannot be equal", nil,
			zap.Int32("position_0", params.Tasks[0].Position),
			zap.Int32("position_1", params.Tasks[1].Position))

		return taskdomain.ErrInvalidPosition()
	}

	err := tp.taskRepository.UpdatePosition(ctx, params.Tasks)
	if err != nil {
		logger.Error("error on update task position", err,
			zap.String("journey", "updateTaskPosition"))

		return taskdomain.ErrCouldNotUpdateTask(err)
	}

	logger.Info("task position updated succesfully",
		zap.String("journey", "updateTaskposition"))

	return nil
}
