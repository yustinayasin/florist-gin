package response

import (
	"florist-gin/business/carts"
	"florist-gin/business/products"
)

type CartResponse struct {
	Id      int `json:"id"`
	UserId  int `json:"user_id"`
	Product []products.Product
}

func FromUsecase(cart carts.Cart) CartResponse {
	return CartResponse{
		Id:      cart.Id,
		UserId:  cart.UserId,
		Product: cart.Product,
	}
}

func FromUsecaseList(cart []carts.Cart) []CartResponse {
	var cartResponse []CartResponse

	for _, v := range cart {
		cartResponse = append(cartResponse, FromUsecase(v))
	}

	return cartResponse
}
