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

func TestMain(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mock_repository.NewMockTaskRepository(ctrl)
	updateCompletion := taskcase.NewUpdateCompletion(r)

	ctx := context.Background()

	t.Run("Should be able to update task completion", func(t *testing.T) {
		taskId := int32(1)
		isCompleted := false

		r.EXPECT().
			CheckExists(gomock.Any(), taskId).
			Return(true, nil)

		r.EXPECT().
			UpdateCompletion(gomock.Any(), taskId, isCompleted).
			Return(nil)

		err := updateCompletion.Execute(ctx, taskId, isCompleted)
		if err != nil {
			t.Fatalf("no error was expected, got: %v", err)
		}
	})

	t.Run("Should fail if the user does not exists", func(t *testing.T) {
		taskId := int32(1)
		isCompleted := false

		r.EXPECT().
			CheckExists(gomock.Any(), taskId).
			Return(false, nil)

		err := updateCompletion.Execute(ctx, taskId, isCompleted)
		if err == nil {
			t.Fatalf("error was expected, got: %v", err)
		}

		res := resp.AsRestErr(err)

		if res.Status != taskdomain.SttCouldNotFindTask {
			t.Fatalf("status expected %v, got: %v",
				res.Status,
				taskdomain.SttCouldNotFindTask,
			)
		}
	})

	t.Run("Should fail to check if the task exists", func(t *testing.T) {
		taskId := int32(1)
		isCompleted := false

		r.EXPECT().
			CheckExists(gomock.Any(), taskId).
			Return(false, errors.New("internal error"))

		err := updateCompletion.Execute(ctx, taskId, isCompleted)
		if err == nil {
			t.Fatalf("error was expected, got: %v", err)
		}

		res := resp.AsRestErr(err)

		if res.Status != taskdomain.SttInternalRepositoryErr {
			t.Fatalf("status expected %v, got: %v",
				res.Status,
				taskdomain.SttInternalRepositoryErr,
			)
		}
	})

	t.Run("Should fail to update task completion", func(t *testing.T) {
		taskId := int32(1)
		isCompleted := false

		r.EXPECT().
			CheckExists(gomock.Any(), taskId).
			Return(true, nil)

		r.EXPECT().
			UpdateCompletion(gomock.Any(), taskId, isCompleted).
			Return(errors.New("internal error"))

		err := updateCompletion.Execute(ctx, taskId, isCompleted)
		if err == nil {
			t.Fatalf("error was expected, got: %v", err)
		}

		res := resp.AsRestErr(err)

		if res.Status != taskdomain.SttCouldNotUpdateTask {
			t.Fatalf("status expected %v, got: %v",
				res.Status,
				taskdomain.SttCouldNotUpdateTask,
			)
		}
	})
}
