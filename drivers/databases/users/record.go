package users

import "florist-gin/business/users"

type User struct {
	Id          int `gorm:"primaryKey;unique;autoIncrement:true"`
	Name        string
	Email       string `gorm:"unique"`
	Password    string
	Address     string
	PhoneNumber string
	PostalCode  string
}

func (user User) ToUsecase() users.User {
	return users.User{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		Address:     user.Address,
		PhoneNumber: user.PhoneNumber,
		PostalCode:  user.PostalCode,
	}
}

func ToUsecaseList(user []User) []users.User {
	var newUsers []users.User

	for _, v := range user {
		newUsers = append(newUsers, v.ToUsecase())
	}
	return newUsers
}

func FromUsecase(user users.User) User {
	return User{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		Address:     user.Address,
		PhoneNumber: user.PhoneNumber,
		PostalCode:  user.PostalCode,
	}
}