package modulecase_test

import (
	"context"
	"errors"
	"testing"

	modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
	mock_repository "github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository/mocks"
	"github.com/golang/mock/gomock"
)

func TestChangeIconModule(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	moduleRepository := mock_repository.NewMockModuleRepository(ctrl)
	updateIconUseCase := modulecase.NewUpdateModuleIconUseCase(moduleRepository)

	ctx := context.Background()
	moduleId := "12345"

	t.Run("Should be able to change the module icon", func(t *testing.T) {
		moduleRepository.EXPECT().UpdateIcon(
			gomock.Any(), gomock.Any(), gomock.Any(),
		).Return(nil)

		icon := "code"
		err := updateIconUseCase.Execute(ctx, moduleId, icon)

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})

	t.Run("Should fail if the uuid is invalid", func(t *testing.T) {
		invalidId := "abc123"
		icon := "code"

		err := updateIconUseCase.Execute(ctx, invalidId, icon)

		if err == nil {
			t.Fatalf("expected invalid uuid error, got %v", err)
		}

		resterr := resp.AsRestErr(err)

		if resterr.Status != moduledomain.SttInvalidModuleID {
			t.Fatalf("expected status %v, got %v",
				moduledomain.SttInvalidModuleID, resterr.Status,
			)
		}
	})

	t.Run("Should fail if the icon is empty", func(t *testing.T) {
		icon := "  "
		err := updateIconUseCase.Execute(ctx, moduleId, icon)

		if err == nil {
			t.Fatalf("expected icon is empty error, got %v", err)
		}

		resterr := resp.AsRestErr(err)

		if resterr.Status != moduledomain.SttNewIconNameIsEmpty {
			t.Fatalf("expected status %v, got %v",
				moduledomain.SttNewIconNameIsEmpty, resterr.Status,
			)
		}
	})

	t.Run("Should fail if unable to update icon", func(t *testing.T) {
		moduleRepository.EXPECT().UpdateIcon(
			gomock.Any(), gomock.Any(), gomock.Any(),
		).Return(errors.New("update error"))

		icon := "coder"
		err := updateIconUseCase.Execute(ctx, moduleId, icon)

		if err == nil {
			t.Fatalf("expected icon is empty error, got %v", err)
		}

		resterr := resp.AsRestErr(err)

		if resterr.Status != moduledomain.SttCouldNotPersistModule {
			t.Fatalf("expected status %v, got %v",
				moduledomain.SttCouldNotPersistModule, resterr.Status,
			)
		}
	})
}
