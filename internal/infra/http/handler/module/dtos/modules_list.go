package moduledto

import (
	"time"

	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
)

type ModulesResponse struct {
	Modules []ModuleListResponse `json:"modules"`
}

type ModuleListResponse struct {
	Id        int32     `json:"id"`
	Name      string    `json:"name"`
	IsOpen    bool      `json:"is_open"`
	Icon      string    `json:"icon"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewModuleListDTO(
	modules []moduledomain.Module,
) (modulesRes ModulesResponse) {
	var moduleList = make([]ModuleListResponse, 0, len(modules))

	for _, m := range modules {
		moduleList = append(
			moduleList,
			ModuleListResponse{
				Id:        m.GetID(),
				Name:      m.GetName(),
				IsOpen:    m.GetIsOpen(),
				Icon:      m.GetIcon(),
				UpdatedAt: m.GetCreateAt(),
				CreatedAt: m.GetUpdatedAt(),
			},
		)
	}

	return ModulesResponse{Modules: moduleList}
}
