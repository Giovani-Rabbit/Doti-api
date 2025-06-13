package userDomain

import "time"

type IUserDomain interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetID() string
}

func NewUserDomain(
	id string,
	name string,
	email string,
	password string,
	createdAt time.Time,
) IUserDomain {
	return &UserDomain{
		id:        id,
		name:      name,
		email:     email,
		password:  password,
		createdAt: createdAt,
	}
}

func NewCreateUserDomain(
	name string, email string, password string,
) IUserDomain {
	return &UserDomain{
		name:     name,
		email:    email,
		password: encryptPassword(password),
	}
}
