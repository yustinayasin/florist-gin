package request

import (
	"florist-gin/business/ordersproducts"
)

type OrdersProducts struct {
	Quantity  int `form:"quantity"`
	Price     int `form:"price"`
	OrderId   int `form:"orderId"`
	ProductId int `form:"productId"`
}

func (ordersProducts *OrdersProducts) ToUsecase() *ordersproducts.OrdersProducts {
	return &ordersproducts.OrdersProducts{
		Quantity:  ordersProducts.Quantity,
		Price:     ordersProducts.Price,
		OrderId:   ordersProducts.OrderId,
		ProductId: ordersProducts.ProductId,
	}
}
