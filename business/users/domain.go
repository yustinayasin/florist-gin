package users

import "florist-gin/business/types"

type User struct {
	Id          int
	Name        string
	Email       string
	Password    string
	Token       string
	Address     string
	PhoneNumber string
	PostalCode  string
	TypeId      int
	Type        types.Type
}

type UserUseCaseInterface interface {
	SignUp(user User) (User, error)
	Login(user User) (User, error)
	EditUser(user User, id int) (User, error)
	DeleteUser(id int) (User, error)
	GetUser(id int) (User, error)
}

type UserRepoInterface interface {
	SignUp(user User) (User, error)
	Login(user User) (User, error)
	EditUser(user User, id int) (User, error)
	DeleteUser(id int) (User, error)
	GetUser(id int) (User, error)
}
