package taskdomain

import "time"

type Task interface {
	ID() int32
	ModuleId() int32
	Name() string
	IsCompleted() bool
	Position() int32
	CreatedAt() time.Time
	UpdatedAt() time.Time

	IsValidToCreate() bool
}

func NewFromDB(
	id int32,
	ModuleId int32,
	name string,
	isCompleted bool,
	position int32,
	createdAt time.Time,
	updatedAt time.Time,
) Task {
	return &taskDomain{
		id:          id,
		moduleId:    ModuleId,
		name:        name,
		isCompleted: isCompleted,
		position:    position,
		createdAt:   createdAt,
		updatedAt:   createdAt,
	}
}

func New(
	moduleId int,
	name string,
	position int,
) Task {
	now := time.Now()
	return &taskDomain{
		moduleId:    int32(moduleId),
		name:        name,
		isCompleted: false,
		position:    int32(position),
		createdAt:   now,
		updatedAt:   now,
	}
}
