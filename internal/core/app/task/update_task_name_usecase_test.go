package taskcase_test

import (
	"context"
	"testing"

	taskcase "github.com/Giovani-Coelho/Doti-API/internal/core/app/task"
	taskdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
	mock_repository "github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository/mocks"
	"github.com/golang/mock/gomock"
)

func TestUpdateTaskName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_repository.NewMockTaskRepository(ctrl)
	updateTaskName := taskcase.NewUpdateTaskName(repo)

	ctx := context.Background()

	t.Run("Should be able to rename task", func(t *testing.T) {
		taskId := int32(123)
		newName := "Task"

		repo.EXPECT().
			UpdateName(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil)

		err := updateTaskName.Execute(ctx, taskId, newName)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})

	t.Run("Should fail with an invalid task name", func(t *testing.T) {
		taskId := int32(123)
		newName := "  "

		err := updateTaskName.Execute(ctx, taskId, newName)
		if err == nil {
			t.Fatalf("expected error, got %v", err)
		}

		res := resp.AsRestErr(err)

		if res.Status != taskdomain.SttInvalidTaskName {
			t.Fatalf("expected status %v, got: %v",
				taskdomain.SttInvalidTaskName,
				res.Status,
			)
		}
	})

	t.Run("Should fail if the task is not found", func(t *testing.T) {
		taskId := int32(123)
		newName := "Task"

		repo.EXPECT().
			UpdateName(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(taskdomain.ErrCouldNotFindTask())

		err := updateTaskName.Execute(ctx, taskId, newName)
		if err == nil {
			t.Fatalf("expected error, got %v", err)
		}

		res := resp.AsRestErr(err)

		if res.Status != taskdomain.SttCouldNotFindTask {
			t.Fatalf("expected status %v, got: %v",
				taskdomain.SttCouldNotFindTask,
				res.Status,
			)
		}
	})
}
