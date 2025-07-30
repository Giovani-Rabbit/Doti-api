package moduledomain

import "time"

type Module interface {
	GetID() string
	GetUserId() string
	GetName() string
	GetIsOpen() bool
	GetIcon() string
	GetCreateAt() time.Time
	GetUpdatedAt() time.Time
}

func New(
	id string,
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
	return &moduleDomain{
		userId: userId,
		name:   name,
		icon:   icon,
	}
}
