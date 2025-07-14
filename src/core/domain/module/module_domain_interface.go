package moduledomain

import "time"

type IModuleDomain interface {
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
) IModuleDomain {
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
