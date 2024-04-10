package response

import (
	"florist-gin/business/cartsproducts"
)

type CartsProductsResponse struct {
	Id        int `json:"id"`
	CartId    int `json:"cartId"`
	ProductId int `json:"productId"`
	Quantity  int `json:"quantity"`
}

func FromUsecase(cartsProducts cartsproducts.CartsProducts) CartsProductsResponse {
	return CartsProductsResponse{
		Id:        cartsProducts.Id,
		CartId:    cartsProducts.CartId,
		ProductId: cartsProducts.ProductId,
		Quantity:  cartsProducts.Quantity,
	}
}

func FromUsecaseList(cartsProduct []cartsproducts.CartsProducts) []CartsProductsResponse {
	var cartResponse []CartsProductsResponse

	for _, v := range cartsProduct {
		cartResponse = append(cartResponse, FromUsecase(v))
	}

	return cartResponse
}
