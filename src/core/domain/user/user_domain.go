package userDomain

import "time"

type UserDomain struct {
	id        string
	email     string
	name      string
	password  string
	createdAt time.Time
	UpdatedAt time.Time
}

func (u *UserDomain) GetID() string       { return u.id }
func (u *UserDomain) GetEmail() string    { return u.email }
func (u *UserDomain) GetName() string     { return u.name }
func (u *UserDomain) GetPassword() string { return u.password }
