package mapper

import (
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/db/sqlc"
)

func ConvertCreateModuleRowToModule(
	module *sqlc.Module,
) (domain moduledomain.Module) {
	return moduledomain.New(
		module.ID.String(),
		module.UserID.String(),
		module.Name,
		module.IsOpen,
		module.Icon,
		module.CreatedAt,
		module.UpdatedAt,
	)
}

func ConvertListModuleByUserIdRowToModules(
	modules *[]sqlc.Module,
) (moduleList []moduledomain.Module) {
	for _, m := range *modules {
		moduleList = append(moduleList, moduledomain.New(
			m.ID.String(),
			m.UserID.String(),
			m.Name,
			m.IsOpen,
			m.Icon,
			m.CreatedAt,
			m.UpdatedAt,
		))
	}
	return
}
