package taskcase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestChangeTaskPosition(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// tr := mock_repository.NewMockTaskRepository(ctrl)
	// changeTaskPosition := taskcase.NewChangeTaskPosition(tr)

	// ctx := context.Background()

	t.Run("Should be able to change the task position", func(t *testing.T) {

	})
}
