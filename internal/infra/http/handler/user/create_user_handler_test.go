package userhandler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_usercase "github.com/Giovani-Coelho/Doti-API/internal/core/app/user/mocks"
	authdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/auth"
	userdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/user"
	userhandler "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/user"
	userdto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/user/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
	"github.com/Giovani-Coelho/Doti-API/internal/pkg/auth"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestCreateUserHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	createUserUseCase := mock_usercase.NewMockCreate(ctrl)
	createUserHandler := userhandler.NewCreateHandler(createUserUseCase)

	authUser := &authdomain.AuthClaims{
		ID:    uuid.NewString(),
		Name:  "Giovani",
		Email: "giovani@example.com",
	}

	t.Run("Should be able to create user", func(t *testing.T) {
		createUserRequest := userdto.CreateUserDTO{
			Name:     "giovani",
			Email:    "giovani@email.com",
			Password: "1234",
		}

		jsonBody, err := json.Marshal(createUserRequest)
		if err != nil {
			t.Fatalf("failed to marshal body: %v", err)
		}

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(jsonBody))

		ctx := context.WithValue(req.Context(), auth.AuthenticatedUserKey, authUser)
		req = req.WithContext(ctx)

		createdResponse := userdomain.NewCreateUser(
			createUserRequest.Name,
			createUserRequest.Email,
			createUserRequest.Password,
		)

		createUserUseCase.EXPECT().
			Execute(gomock.Any(), gomock.Any()).
			Return(createdResponse, nil)

		rr := httptest.NewRecorder()
		createUserHandler.Execute(rr, req)

		body, err := io.ReadAll(rr.Body)
		if err != nil {
			t.Fatalf("failed to get body: %v", err)
		}

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status %d, got %d body: %s",
				http.StatusCreated, rr.Code, body,
			)
		}

		var res userdto.UserCreatedResponse
		err = json.Unmarshal(body, &res)
		if err != nil {
			t.Fatalf("failed to unmarshal body: %v", err)
		}
	})

	t.Run("Should fail if the body is invalid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/users", nil)

		ctx := context.WithValue(req.Context(), auth.AuthenticatedUserKey, authUser)
		req = req.WithContext(ctx)

		rr := httptest.NewRecorder()
		createUserHandler.Execute(rr, req)

		body, err := io.ReadAll(rr.Body)
		if err != nil {
			t.Fatalf("failed to get response body: %v", err)
		}

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected code %d, got %d body: %s",
				http.StatusCreated, rr.Code, body,
			)
		}

		var res resp.RestErr
		err = json.Unmarshal(body, &res)
		if err != nil {
			t.Fatalf("failed to unmarshal body: %v", err)
		}

		if res.Status != resp.SttInvalidRequestBody {
			t.Errorf("expected status %s, got %s",
				resp.SttInvalidRequestBody, res.Status,
			)
		}
	})
}
