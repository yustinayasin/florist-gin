package response

import (
	"florist-gin/business/ordersproducts"
	"florist-gin/business/products"
	"time"
)

type OrdersProductsResponse struct {
	Id             int
	Quantity       int
	Price          int
	OrderId        int
	ProductId      int
	OrdersProducts ordersproducts.OrdersProducts
	Product        products.Product
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func FromUsecase(ordersProducts ordersproducts.OrdersProducts) OrdersProductsResponse {
	return OrdersProductsResponse{
		Id:        ordersProducts.Id,
		Quantity:  ordersProducts.Quantity,
		Price:     ordersProducts.Price,
		OrderId:   ordersProducts.OrderId,
		ProductId: ordersProducts.ProductId,
		CreatedAt: ordersProducts.CreatedAt,
		UpdatedAt: ordersProducts.UpdatedAt,
	}
}

func FromUsecaseList(ordersProducts []ordersproducts.OrdersProducts) []OrdersProductsResponse {
	var orderResponse []OrdersProductsResponse

	for _, v := range ordersProducts {
		orderResponse = append(orderResponse, FromUsecase(v))
	}

	return orderResponse
}
