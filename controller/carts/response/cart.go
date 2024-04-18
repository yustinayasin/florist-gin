package response

import (
	"florist-gin/business/carts"
	"florist-gin/business/products"
	"time"
)

type CartResponse struct {
	Id        int `form:"id"`
	UserId    int `form:"user_id"`
	Product   []products.Product
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FromUsecase(cart carts.Cart) CartResponse {
	return CartResponse{
		Id:        cart.Id,
		UserId:    cart.UserId,
		Product:   cart.Product,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
	}
}

func FromUsecaseList(cart []carts.Cart) []CartResponse {
	var cartResponse []CartResponse

	for _, v := range cart {
		cartResponse = append(cartResponse, FromUsecase(v))
	}

	return cartResponse
}
