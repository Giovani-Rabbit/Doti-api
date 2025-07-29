package mapper

import (
	"time"

	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/db/sqlc"
)

func FromCreateModuleRow(
	module *sqlc.CreateModuleRow,
) (domain moduledomain.IModuleDomain) {
	return moduledomain.New(
		module.ID.String(),
		"",
		module.Name,
		module.IsOpen,
		module.Icon,
		time.Now(),
		time.Now(),
	)
}
