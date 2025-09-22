package taskcase

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
)

type ChangeTaskPosition interface {
	Execute(ctx context.Context, taskId, position int) error
}

type changeTaskPosition struct {
	taskRepository repository.TaskRepository
}

func NewChangeTaskPosition(
	tr repository.TaskRepository,
) ChangeTaskPosition {
	return &changeTaskPosition{
		taskRepository: tr,
	}
}

func (tp *changeTaskPosition) Execute(
	ctx context.Context, taskId, position int,
) error {

	return nil
}
