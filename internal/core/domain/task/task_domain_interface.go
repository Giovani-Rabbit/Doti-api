package taskdomain

import "time"

type TaskDomain interface {
	ID() int32
	ModuleID() int32
	Name() string
	IsComplete() bool
	Position() int
	CreatedAt() time.Time
	UpdatedAt() time.Time
}

func NewTask(id, moduleID int32, name string, position int) TaskDomain {
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
		id:         id,
		moduleID:   moduleID,
		name:       name,
		isComplete: false,
		position:   position,
		createdAt:  now,
		updatedAt:  now,
	}
}
