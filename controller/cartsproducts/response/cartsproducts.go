package response

import (
	"florist-gin/business/cartsproducts"
	"time"
)

type CartsProductsResponse struct {
	Id        uint32 `json:"id"`
	CartId    uint32 `json:"cartId"`
	ProductId uint32 `json:"productId"`
	Quantity  int    `json:"quantity"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FromUsecase(cartsProducts cartsproducts.CartsProducts) CartsProductsResponse {
	return CartsProductsResponse{
		Id:        cartsProducts.Id,
		CartId:    cartsProducts.CartId,
		ProductId: cartsProducts.ProductId,
		Quantity:  cartsProducts.Quantity,
		CreatedAt: cartsProducts.CreatedAt,
		UpdatedAt: cartsProducts.UpdatedAt,
	}
}

func FromUsecaseList(cartsProduct []cartsproducts.CartsProducts) []CartsProductsResponse {
	var cartResponse []CartsProductsResponse

	for _, v := range cartsProduct {
		cartResponse = append(cartResponse, FromUsecase(v))
	}

	return cartResponse
}
