package response

import (
	"florist-gin/business/carts"
	"florist-gin/controller/products/response"
	"time"
)

type CartResponse struct {
	Id        int `form:"id"`
	UserId    int `form:"user_id"`
	Product   []response.ProductResponse
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FromUsecase(cart carts.Cart) CartResponse {
	newProducts := response.FromUsecaseList(cart.Product)

	return CartResponse{
		Id:        cart.Id,
		UserId:    cart.UserId,
		Product:   newProducts,
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
