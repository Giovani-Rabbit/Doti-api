package moduledto

import (
	"time"

	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
)

type CreateModuleDTO struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type CreateModuleResponse struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	IsOpen    bool      `json:"is_open"`
	Icon      string    `json:"icon"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewModuleCreatedResponse(module moduledomain.Module) CreateModuleResponse {
	return CreateModuleResponse{
		ID:        module.ID(),
		Name:      module.Name(),
		IsOpen:    module.IsOpen(),
		Icon:      module.Icon(),
		CreatedAt: module.CreateAt(),
		UpdatedAt: module.UpdatedAt(),
	}
}
