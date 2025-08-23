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
	"github.com/google/uuid"
)

func TestChangeIconModule(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	moduleRepository := mock_repository.NewMockModuleRepository(ctrl)
	ChangeIconUseCase := modulecase.NewChangeModuleIconUseCase(moduleRepository)

	ctx := context.Background()
	moduleId := uuid.NewString()

	t.Run("Should be able to change the module icon", func(t *testing.T) {
		moduleRepository.EXPECT().UpdateIcon(
			gomock.Any(), gomock.Any(), gomock.Any(),
		).Return(nil)

		icon := "code"
		err := ChangeIconUseCase.Execute(ctx, moduleId, icon)

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})

	t.Run("Should fail if the uuid is invalid", func(t *testing.T) {
		invalidId := "123"
		icon := "code"

		err := ChangeIconUseCase.Execute(ctx, invalidId, icon)

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
		err := ChangeIconUseCase.Execute(ctx, moduleId, icon)

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
		err := ChangeIconUseCase.Execute(ctx, moduleId, icon)

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
