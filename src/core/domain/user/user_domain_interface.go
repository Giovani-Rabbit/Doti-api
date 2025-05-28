package userDomain

type IUserDomain interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetID() string
}

func NewUserDomain(
	email string, name string, password string,
) IUserDomain {
	return &userDomain{
		name:     name,
		email:    email,
		password: encryptPassword(password),
	}
}
