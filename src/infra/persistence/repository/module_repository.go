package repository

import "context"

type IModuleRepository interface {
	Create(ctx context.Context)
}
