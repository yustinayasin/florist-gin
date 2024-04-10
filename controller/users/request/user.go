package request

import "florist-gin/business/users"

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserEdit struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	PostalCode  string `json:"postalCode"`
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