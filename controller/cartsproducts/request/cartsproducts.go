package request

import (
	"florist-gin/business/cartsproducts"
)

type CartsProducts struct {
	CartId    uint32 `json:"cartId"`
	ProductId uint32 `json:"productId"`
	Quantity  int    `json:"quantity"`
}

func (cartsProducts *CartsProducts) ToUsecase() *cartsproducts.CartsProducts {
	return &cartsproducts.CartsProducts{
		CartId:    cartsProducts.CartId,
		ProductId: cartsProducts.ProductId,
		Quantity:  cartsProducts.Quantity,
	}
}
