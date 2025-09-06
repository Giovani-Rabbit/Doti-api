package moduledomain

import "time"

type Module interface {
	GetID() int32
	GetUserId() string
	GetName() string
	GetIsOpen() bool
	GetIcon() string
	GetCreateAt() time.Time
	GetUpdatedAt() time.Time

	IsValid() bool
}

func New(
	id int32,
	user_id string,
	name string,
	isOpen bool,
	icon string,
	createdAt time.Time,
	updatedAt time.Time,
) Module {
	return &moduleDomain{
		id:        id,
		userId:    user_id,
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
