package user_entity

type User struct {
	Id   string
	Name string
}

var Users []User

type IUserRepository interface {
	createUserRepository()
}
