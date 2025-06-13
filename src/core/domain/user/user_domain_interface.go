package userDomain

import "time"

type IUserDomain interface {
	GetID() string
	GetName() string
	GetEmail() string
	GetPassword() string
	GetCreateAt() time.Time
	GetUpdatedAt() time.Time
}

func NewUserDomain(
	id string,
	name string,
	email string,
	password string,
	createdAt time.Time,
	updatedAt time.Time,
) IUserDomain {
	return &userDomain{
		id:        id,
		name:      name,
		email:     email,
		password:  password,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func NewCreateUserDomain(
	name string, email string, password string,
) IUserDomain {
	return &userDomain{
		name:     name,
		email:    email,
		password: encryptPassword(password),
	}
}
