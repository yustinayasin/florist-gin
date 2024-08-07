package request

import (
	"florist-gin/business/cartsproducts"
)

type CartsProducts struct {
	CartId    int `form:"cartId"`
	ProductId int `form:"productId"`
	Quantity  int `form:"quantity"`
}

func (cartsProducts *CartsProducts) ToUsecase() *cartsproducts.CartsProducts {
	return &cartsproducts.CartsProducts{
		CartId:    cartsProducts.CartId,
		ProductId: cartsProducts.ProductId,
		Quantity:  cartsProducts.Quantity,
	}
}
