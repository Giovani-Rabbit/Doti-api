package moduledomain

import "time"

type moduleDomain struct {
	id        string
	userId    string
	name      string
	isOpen    bool
	icon      string
	createdAt time.Time
	updatedAt time.Time
}

func (md *moduleDomain) GetID() string           { return md.id }
func (md *moduleDomain) GetUserId() string       { return md.userId }
func (md *moduleDomain) GetName() string         { return md.name }
func (md *moduleDomain) GetIsOpen() bool         { return md.isOpen }
func (md *moduleDomain) GetIcon() string         { return md.icon }
func (md *moduleDomain) GetCreateAt() time.Time  { return md.createdAt }
func (md *moduleDomain) GetUpdatedAt() time.Time { return md.updatedAt }
