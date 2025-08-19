package modulehandler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module/mocks"
	authdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/auth"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	modulehandler "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/module"
	moduledto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/module/dtos"
	"github.com/Giovani-Coelho/Doti-API/internal/pkg/auth"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestCreateModuleHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	createModuleUseCase := mock_modulecase.NewMockCreate(ctrl)
	createModuleHandler := modulehandler.NewCreateModuleHandler(createModuleUseCase)

	authUser := &authdomain.AuthClaims{
		ID:    uuid.NewString(),
		Name:  "Giovani",
		Email: "giovani@example.com",
	}

	t.Run("Should be able to create a module", func(t *testing.T) {
		createModule := moduledto.CreateModuleDTO{
			Name: "new Module",
			Icon: "Icon",
		}

		jsonBody, err := json.Marshal(createModule)
		if err != nil {
			t.Fatalf("failed to marshal body: %v", err)
		}

		req := httptest.NewRequest(http.MethodPost, "/modules", bytes.NewReader(jsonBody))

		ctx := context.WithValue(req.Context(), auth.AuthenticatedUserKey, authUser)
		req = req.WithContext(ctx)

		createdModule := moduledomain.NewCreateModule(
			uuid.NewString(), createModule.Name, createModule.Icon,
		)

		createModuleUseCase.EXPECT().
			Execute(gomock.Any(), gomock.Any()).
			Return(createdModule, nil)

		rr := httptest.NewRecorder()
		createModuleHandler.Execute(rr, req)

		body, err := io.ReadAll(rr.Body)
		if err != nil {
			t.Fatalf("failed to get body: %v", err)
		}

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status %d, got %d body: %s",
				http.StatusCreated, rr.Code, body,
			)
		}

		var res moduledto.CreateModuleResponse
		err = json.Unmarshal(body, &res)
		if err != nil {
			t.Fatalf("invalid return. Failed to unmarshal body: %v", err)
		}
	})
}
