package taskdomain

import "time"

type Task interface {
	ID() int32
	ModuleID() int32
	Name() string
	IsCompleted() bool
	Position() int32
	CreatedAt() time.Time
	UpdatedAt() time.Time
}

func NewFromDB(
	id int32,
	moduleID int32,
	name string,
	isCompleted bool,
	position int32,
	createdAt time.Time,
	updatedAt time.Time,
) Task {
	return &taskDomain{
		id:          id,
		moduleID:    moduleID,
		name:        name,
		isCompleted: isCompleted,
		position:    position,
		createdAt:   createdAt,
		updatedAt:   createdAt,
	}
}

func New(
	id, moduleID int32,
	name string,
	position int,
) Task {
	if id <= 0 {
		return nil
	}

	if moduleID <= 0 {
		return nil
	}

	if name == "" {
		return nil
	}

	now := time.Now()
	return &taskDomain{
		id:          id,
		moduleID:    moduleID,
		name:        name,
		isCompleted: false,
		position:    int32(position),
		createdAt:   now,
		updatedAt:   now,
	}
}
