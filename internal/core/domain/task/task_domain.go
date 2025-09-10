package taskdomain

import "time"

type taskDomain struct {
	id         int32
	moduleID   int32
	name       string
	isComplete bool
	position   int
	createdAt  time.Time
	updatedAt  time.Time
}

func (t *taskDomain) ID() int32            { return t.id }
func (t *taskDomain) ModuleID() int32      { return t.moduleID }
func (t *taskDomain) Name() string         { return t.name }
func (t *taskDomain) IsComplete() bool     { return t.isComplete }
func (t *taskDomain) Position() int        { return t.position }
func (t *taskDomain) CreatedAt() time.Time { return t.createdAt }
func (t *taskDomain) UpdatedAt() time.Time { return t.updatedAt }
