package detailscase_test

import (
	"context"
	"testing"

	detailscase "github.com/Giovani-Coelho/Doti-API/internal/core/app/task_details"
	taskdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
	mock_repository "github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository/mocks"
	"github.com/golang/mock/gomock"
)

func TestTaskDetailsUpdateDescription(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_repository.NewMockTaskDetailsRepository(ctrl)
	usecase := detailscase.NewTaskDetailsUseCase(repo)

	ctx := context.Background()

	t.Run("Should be able to update the description", func(t *testing.T) {
		taskId := int32(12345)
		description := "blablabla"

		repo.EXPECT().
			UpdateDescription(ctx, taskId, description).
			Return(nil)

		err := usecase.Execute(ctx, taskId, description)
		if err != nil {
			t.Errorf("expected no error, got: %v", err)
		}
	})

	t.Run("Should fail if task details are not found", func(t *testing.T) {
		taskId := int32(12345)
		description := "blablabla"

		repo.EXPECT().
			UpdateDescription(ctx, taskId, description).
			Return(taskdomain.ErrCouldNotFindTask())

		err := usecase.Execute(ctx, taskId, description)
		if err == nil {
			t.Errorf("expected error, got: %v", err)
		}

		respErr := resp.AsRestErr(err)

		if respErr.Status != taskdomain.SttCouldNotFindTask {
			t.Errorf("expected status %v, got: %v",
				taskdomain.SttCouldNotFindTask,
				respErr.Status,
			)
		}
	})
}
