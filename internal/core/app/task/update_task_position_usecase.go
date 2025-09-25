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

	if taskdomain.HasRepeatedPositions(params.MovedTasks) {
		logger.Error("has repeated position", nil,
			zap.String("journey", "updateTaskPosition"))

		return taskdomain.ErrRepeatedPosition()
	}

	err := tp.taskRepository.UpdatePosition(ctx, params.MovedTasks)
	if err != nil {
		logger.Error("error on update task position", err,
			zap.String("journey", "updateTaskPosition"))

		return taskdomain.ErrCouldNotUpdateTask(err)
	}

	logger.Info("task position updated succesfully",
		zap.String("journey", "updateTaskposition"))

	return nil
}
