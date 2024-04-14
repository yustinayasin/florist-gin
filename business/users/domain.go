package users

import (
	"florist-gin/business/types"
	"time"
)

type User struct {
	Id          uint32
	Name        string
	Email       string
	Password    string
	Token       string
	Address     string
	PhoneNumber string
	PostalCode  string
	TypeId      uint32
	Type        types.Type
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserUseCaseInterface interface {
	SignUp(user User) (User, error)
	Login(user User) (User, error)
	EditUser(user User, id uint32) (User, error)
	DeleteUser(id uint32) (User, error)
	GetUser(id uint32) (User, error)
}

type UserRepoInterface interface {
	SignUp(user User) (User, error)
	Login(user User) (User, error)
	EditUser(user User, id uint32) (User, error)
	DeleteUser(id uint32) (User, error)
	GetUser(id uint32) (User, error)
}
