package userDomain

import "time"

type userDomain struct {
	id        string
	email     string
	name      string
	password  string
	createdAt time.Time
	UpdatedAt time.Time
}

func (u *userDomain) GetID() string       { return u.id }
func (u *userDomain) GetEmail() string    { return u.email }
func (u *userDomain) GetName() string     { return u.name }
func (u *userDomain) GetPassword() string { return u.password }
