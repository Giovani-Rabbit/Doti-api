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
	userdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/user"
	modulehandler "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/module"
	moduledto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/module/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
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
		rr := httptest.NewRecorder()

		createdModule := moduledomain.NewCreateModule(
			uuid.NewString(), createModule.Name, createModule.Icon,
		)

		createModuleUseCase.EXPECT().
			Execute(gomock.Any(), gomock.Any()).
			Return(createdModule, nil)

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

	t.Run("Should fail if not logged in", func(t *testing.T) {
		createModule := moduledto.CreateModuleDTO{
			Name: "new Module",
			Icon: "Icon",
		}

		jsonBody, err := json.Marshal(createModule)
		if err != nil {
			t.Fatalf("failed to marshal body: %v", err)
		}

		req := httptest.NewRequest(http.MethodPost, "/modules", bytes.NewReader(jsonBody))
		rr := httptest.NewRecorder()

		createModuleHandler.Execute(rr, req)

		body, err := io.ReadAll(rr.Body)
		if err != nil {
			t.Fatalf("failed to get body: %v", err)
		}

		if rr.Code != http.StatusUnauthorized {
			t.Errorf("expected status %d, got %d body: %s",
				http.StatusUnauthorized, rr.Code, body,
			)
		}

		var res *resp.RestErr
		err = json.Unmarshal(body, &res)
		if err != nil {
			t.Fatalf("invalid return. Failed to unmarshal body: %v", err)
		}

		if res.Status != userdomain.SttUserUnauthorized {
			t.Fatalf("expected status %s, got: %v", userdomain.SttCouldNotFindUser, res.Status)
		}
	})
}
