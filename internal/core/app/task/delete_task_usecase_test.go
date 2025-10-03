package taskcase_test

import (
	"context"
	"errors"
	"testing"

	taskcase "github.com/Giovani-Coelho/Doti-API/internal/core/app/task"
	taskdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
	mock_repository "github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository/mocks"
	"github.com/golang/mock/gomock"
)

func TestDeleteTaskUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock_repository.NewMockTaskRepository(ctrl)
	usecase := taskcase.NewDeleteTaskUseCase(repository)

	ctx := context.Background()

	t.Run("Should be able to delete a task", func(t *testing.T) {
		taskId := int32(32)

		repository.EXPECT().
			Delete(ctx, taskId).
			Return(nil)

		err := usecase.Execute(ctx, taskId)
		if err != nil {
			t.Fatalf("error was not expected, got: %v", err)
		}
	})

	t.Run("Expect an error on delete task", func(t *testing.T) {
		taskId := int32(32)
		repoErr := errors.New("internal error")

		repository.EXPECT().
			Delete(ctx, taskId).
			Return(repoErr)

		err := usecase.Execute(ctx, taskId)
		if err == nil {
			t.Fatalf("error was expected, got: %v", err)
		}

		res := resp.AsRestErr(err)

		if res.Status != taskdomain.SttCouldNotDeleteTask {
			t.Fatalf("error expected: %v, got: %v",
				taskdomain.SttCouldNotDeleteTask,
				res.Status,
			)
		}
	})
}
