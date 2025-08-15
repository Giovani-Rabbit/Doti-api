package modulehandler_test

import (
	"testing"
)

func TestCreateModuleHandler(t *testing.T) {
	// ctrl := gomock.NewController(t)
	// defer ctrl.Finish()

	// moduleRepo := mock_repository.NewMockModuleRepository(ctrl)

	// createModuleCase := modulecase.NewCreateModuleUseCase(moduleRepo)
	// deleteModuleCase := modulecase.NewDeleteModuleUseCase(moduleRepo)
	// getModuleCase := modulecase.NewGetModulesUseCase(moduleRepo)
	// renameModuleCase := modulecase.NewRenameModuleUseCase(moduleRepo)

	// moduleHandler := modulehandler.New(
	// 	createModuleCase,
	// 	getModuleCase,
	// 	renameModuleCase,
	// 	deleteModuleCase,
	// )

	// authUser := &authdomain.AuthClaims{
	// 	ID:    uuid.NewString(),
	// 	Name:  "Giovani",
	// 	Email: "giovani@example.com",
	// }

	// createdModule := moduledomain.New(
	// 	uuid.NewString(),
	// 	authUser.ID,
	// 	authUser.Name,
	// 	false,
	// 	"Icon",
	// 	time.Now(),
	// 	time.Now(),
	// )

	// t.Run("Should be able to create a module", func(t *testing.T) {
	// 	createModule := moduledto.CreateModuleDTO{
	// 		Name: "giovani",
	// 		Icon: "Icon",
	// 	}

	// 	jsonBody, err := json.Marshal(createModule)
	// 	if err != nil {
	// 		t.Fatalf("failed to marshal body: %v", err)
	// 	}

	// 	req := httptest.NewRequest(http.MethodPost, "/module", bytes.NewReader(jsonBody))

	// 	ctx := context.WithValue(req.Context(), auth.AuthenticatedUserKey, authUser)
	// 	req = req.WithContext(ctx)

	// 	moduleRepo.EXPECT().
	// 		Create(gomock.Any(), gomock.Any()).
	// 		Return(createdModule, nil)

	// 	rr := httptest.NewRecorder()
	// 	moduleHandler.CreateModule(rr, req)

	// 	body, err := io.ReadAll(rr.Body)
	// 	if err != nil {
	// 		t.Fatalf("failed to get body: %v", err)
	// 	}

	// 	if rr.Code != http.StatusCreated {
	// 		t.Errorf("expected status %d, got %d body: %s", http.StatusOK, rr.Code, body)
	// 	}

	// 	var res moduledto.CreateModuleResponse
	// 	err = json.Unmarshal(body, &res)
	// 	if err != nil {
	// 		t.Fatalf("failed to unmarshal body: %v", err)
	// 	}

	// 	if res.ID != createdModule.GetID() {
	// 		t.Fatalf("Error: want %s got %s", res.ID, createdModule.GetID())
	// 	}

	// 	if !res.CreatedAt.Equal(createdModule.GetCreateAt()) {
	// 		t.Fatalf("Error: want %s got %s", res.CreatedAt, createdModule.GetCreateAt())
	// 	}
	// })
}
