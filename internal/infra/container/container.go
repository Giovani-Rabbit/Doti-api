package container

import (
	"database/sql"
)

type container struct {
	Module *ModuleHandler
	User   *UserHandler
}

func Setup(db *sql.DB) *container {
	repo := newRepository(db)

	moduleCase := newModuleCase(repo.moduleRepo)
	userCase := newUserCase(repo.userRepo)

	return &container{
		Module: newModuleHandler(moduleCase),
		User:   newUserHandler(userCase),
	}
}
