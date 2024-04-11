package users

import (
	"florist-gin/business/users"
	types "florist-gin/drivers/databases/type"
)

type User struct {
	Id          int `gorm:"primaryKey;unique;autoIncrement:true"`
	Name        string
	Email       string `gorm:"unique"`
	Password    string
	Address     string
	PhoneNumber string
	PostalCode  string
	TypeId      int
	Type        types.Type `gorm:"foreignKey:TypeId"`
}

func (user User) ToUsecase() users.User {
	newType := types.Type.ToUseCase(user.Type)

	return users.User{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		Address:     user.Address,
		PhoneNumber: user.PhoneNumber,
		PostalCode:  user.PostalCode,
		TypeId:      user.TypeId,
		Type:        newType,
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
	newType := types.FromUsecase(user.Type)

	return User{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		Address:     user.Address,
		PhoneNumber: user.PhoneNumber,
		PostalCode:  user.PostalCode,
		TypeId:      user.TypeId,
		Type:        newType,
	}
}
