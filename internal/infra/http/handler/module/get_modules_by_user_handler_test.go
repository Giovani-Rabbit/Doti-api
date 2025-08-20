package modulehandler_test

import (
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

func TestGetModulesByUserHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	getModulesByUserUseCase := mock_modulecase.NewMockGetByUser(ctrl)
	getModulesByUserHandler := modulehandler.NewGetHandler(getModulesByUserUseCase)

	authUser := &authdomain.AuthClaims{
		ID:    uuid.NewString(),
		Name:  "Giovani",
		Email: "giovani@example.com",
	}

	t.Run("Should be able to get modules by user", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/modules", nil)

		ctx := context.WithValue(req.Context(), auth.AuthenticatedUserKey, authUser)

		req = req.WithContext(ctx)
		rr := httptest.NewRecorder()

		var modules []moduledomain.Module
		for range 10 {
			module := moduledomain.NewCreateModule(authUser.ID, "new module", "Icon")
			modules = append(modules, module)
		}

		getModulesByUserUseCase.EXPECT().
			Execute(gomock.Any(), gomock.Any()).
			Return(modules, nil)

		getModulesByUserHandler.Execute(rr, req)

		body, err := io.ReadAll(rr.Body)
		if err != nil {
			t.Fatalf("failed to get body: %v", err)
		}

		if rr.Code != http.StatusOK {
			t.Errorf("expected status %d, got %d body: %s",
				http.StatusOK, rr.Code, body,
			)
		}

		var res moduledto.ModulesResponse
		err = json.Unmarshal(body, &res)
		if err != nil {
			t.Fatalf("invalid return. Failed to unmarshal body: %v", err)
		}

		if len(res.Modules) != 10 {
			t.Fatalf("expected 10 modules, got %d", len(res.Modules))
		}
	})

	t.Run("Should fail if user is not logged", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/modules", nil)
		rr := httptest.NewRecorder()

		getModulesByUserHandler.Execute(rr, req)

		body, err := io.ReadAll(rr.Body)
		if err != nil {
			t.Fatalf("failed to get body: %v", err)
		}

		if rr.Code != http.StatusUnauthorized {
			t.Errorf("expected status %d, got %d body: %s",
				http.StatusOK, rr.Code, body,
			)
		}

		var res *resp.RestErr
		err = json.Unmarshal(body, &res)
		if err != nil {
			t.Fatalf("invalid return. Failed to unmarshal body: %v", err)
		}

		if res.Status != userdomain.SttUserUnauthorized {
			t.Fatalf("expected %v, got %v", userdomain.SttUserUnauthorized, res.Status)
		}
	})
}
