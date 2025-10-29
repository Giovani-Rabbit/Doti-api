package detailscase_test

import (
	"context"
	"testing"

	detailscase "github.com/Giovani-Coelho/Doti-API/internal/core/app/task_details"
	mock_repository "github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository/mocks"
	"github.com/golang/mock/gomock"
)

func TestTaskDetailsUpdatePomodoTarget(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_repository.NewMockTaskDetailsRepository(ctrl)
	usecase := detailscase.NewTaskDetailsUpdatePomodoroTarget(repo)

	ctx := context.Background()

	t.Run("Should be able to update pomodoro target", func(t *testing.T) {
		taskId := int32(12345)
		target := 4

		repo.EXPECT().
			UpdatePomodoroTarget(ctx, taskId, target).
			Return(nil)

		err := usecase.Execute(ctx, taskId, target)
		if err != nil {
			t.Errorf("expected no error, got: %v", err)
		}
	})
}
