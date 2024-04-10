package request

import (
	"florist-gin/business/carts"
)

type Cart struct {
	UserId int `json:"user_id"`
}

func (cart *Cart) ToUsecase() *carts.Cart {
	return &carts.Cart{
		UserId: cart.UserId,
	}
}
