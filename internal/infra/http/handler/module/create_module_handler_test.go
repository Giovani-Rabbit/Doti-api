package modulehandler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module"
	authdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/auth"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	modulehandler "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/module"
	moduledto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/module/dtos"
	mock_repository "github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository/mocks"
	"github.com/Giovani-Coelho/Doti-API/internal/pkg/auth"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestCreateModuleHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	moduleRepo := mock_repository.NewMockModuleRepository(ctrl)

	createModuleCase := modulecase.NewCreateModuleUseCase(moduleRepo)
	deleteModuleCase := modulecase.NewDeleteModuleUseCase(moduleRepo)
	getModuleCase := modulecase.NewGetModulesUseCase(moduleRepo)
	renameModuleCase := modulecase.NewRenameModuleUseCase(moduleRepo)

	moduleHandler := modulehandler.New(
		createModuleCase,
		getModuleCase,
		renameModuleCase,
		deleteModuleCase,
	)

	authUser := &authdomain.AuthClaims{
		ID:    uuid.NewString(),
		Name:  "Giovani",
		Email: "giovani@example.com",
	}

	createdModule := moduledomain.New(
		uuid.NewString(),
		authUser.ID,
		authUser.Name,
		false,
		"Icon",
		time.Now(),
		time.Now(),
	)

	t.Run("Should be able to create a module", func(t *testing.T) {
		createModule := moduledto.CreateModuleDTO{
			Name: "giovani",
			Icon: "Icon",
		}

		jsonBody, err := json.Marshal(createModule)
		if err != nil {
			t.Fatalf("failed to marshal body: %v", err)
		}

		req := httptest.NewRequest(http.MethodPost, "/module", bytes.NewReader(jsonBody))
		rr := httptest.NewRecorder()

		ctx := context.WithValue(req.Context(), auth.AuthenticatedUserKey, authUser)
		req = req.WithContext(ctx)

		moduleRepo.EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(createdModule, nil)

		moduleHandler.CreateModule(rr, req)

		body, err := io.ReadAll(rr.Body)
		if err != nil {
			t.Fatalf("failed to get body: %v", err)
		}

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status %d, got %d body: %s", http.StatusOK, rr.Code, body)
		}

		var resp moduledto.CreateModuleResponse
		err = json.Unmarshal(body, &resp)
		if err != nil {
			t.Fatalf("failed to unmarshal body: %v", err)
		}

		if resp.ID != createdModule.GetID() {
			t.Fatalf("Error: want %s got %s", resp.ID, createdModule.GetID())
		}

		if !resp.CreatedAt.Equal(createdModule.GetCreateAt()) {
			t.Fatalf("Error: want %s got %s", resp.CreatedAt, createdModule.GetCreateAt())
		}
	})
}
