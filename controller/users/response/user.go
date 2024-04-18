package response

import (
	"florist-gin/business/types"
	"florist-gin/business/users"
	"time"
)

type UserResponse struct {
	Id          int        `form:"id"`
	Email       string     `form:"email"`
	Name        string     `form:"name"`
	Password    string     `form:"password"`
	Address     string     `form:"address"`
	PhoneNumber string     `form:"phoneNumber"`
	PostalCode  string     `form:"postalCode"`
	Token       string     `form:"token"`
	TypeId      int        `form:"typeID"`
	Type        types.Type `form:"type"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func FromUsecase(user users.User) UserResponse {
	return UserResponse{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		Address:     user.Address,
		PhoneNumber: user.PhoneNumber,
		PostalCode:  user.PostalCode,
		Token:       user.Token,
		TypeId:      user.TypeId,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}

func FromUsecaseList(user []users.User) []UserResponse {
	var userResponse []UserResponse

	for _, v := range user {
		userResponse = append(userResponse, FromUsecase(v))
	}

	return userResponse
}
