package carts

import (
	"florist-gin/business/carts"
	"time"
)

type Cart struct {
	Id        int `gorm:"primaryKey;unique"`
	UserId    int `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (cart Cart) ToUseCase() carts.Cart {
	return carts.Cart{
		Id:        cart.Id,
		UserId:    cart.UserId,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
	}
}

func FromUsecase(cart carts.Cart) Cart {
	return Cart{
		Id:        cart.Id,
		UserId:    cart.UserId,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
	}
}
