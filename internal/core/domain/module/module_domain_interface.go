package moduledomain

import "time"

type Module interface {
	ID() int32
	UserId() string
	Name() string
	IsOpen() bool
	Icon() string
	CreateAt() time.Time
	UpdatedAt() time.Time

	IsValid() bool
}

func New(
	id int32,
	userId string,
	name string,
	isOpen bool,
	icon string,
	createdAt time.Time,
	updatedAt time.Time,
) Module {
	return &moduleDomain{
		id:        id,
		userId:    userId,
		name:      name,
		isOpen:    isOpen,
		icon:      icon,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func NewCreateModule(
	userId string,
	name string,
	icon string,
) Module {
	now := time.Now()
	return &moduleDomain{
		userId:    userId,
		name:      name,
		icon:      icon,
		isOpen:    false,
		createdAt: now,
		updatedAt: now,
	}
}
