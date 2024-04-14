package response

import (
	"florist-gin/business/carts"
	"florist-gin/business/products"
	"time"
)

type CartResponse struct {
	Id        uint32 `json:"id"`
	UserId    uint32 `json:"user_id"`
	Product   []products.Product
	CreatedAt time.Time
	UpdatedAt time.Time
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
