package request

import "florist-gin/business/users"

type UserLogin struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type UserEdit struct {
	Name        string `form:"name"`
	Email       string `form:"email"`
	Password    string `form:"password"`
	Address     string `form:"address"`
	PhoneNumber string `form:"phoneNumber"`
	PostalCode  string `form:"postalCode"`
}

func (user *UserLogin) ToUsecase() *users.User {
	return &users.User{
		Email:    user.Email,
		Password: user.Password,
	}
}

func (user *UserEdit) ToUsecase() *users.User {
	return &users.User{
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		Address:     user.Address,
		PhoneNumber: user.PhoneNumber,
		PostalCode:  user.PostalCode,
	}
}
