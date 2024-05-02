package ordersproducts

import (
	"florist-gin/business/orders"
	"florist-gin/business/products"
	"time"
)

type OrdersProducts struct {
	Id        int
	Quantity  int
	Price     int
	OrderId   int
	ProductId int
	Order     orders.Order
	Product   products.Product
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrdersProductsUseCaseInterface interface {
	AddOrdersProducts(ordersProducts OrdersProducts) (OrdersProducts, error)
	EditOrdersProducts(ordersProducts OrdersProducts, id int) (OrdersProducts, error)
	DeleteOrdersProducts(id int) (OrdersProducts, error)
	GetOrdersProductsDetail(id int) (OrdersProducts, error)
}

type OrdersProductsRepoInterface interface {
	AddOrdersProducts(ordersProducts OrdersProducts) (OrdersProducts, error)
	EditOrdersProducts(ordersProducts OrdersProducts, id int) (OrdersProducts, error)
	DeleteOrdersProducts(id int) (OrdersProducts, error)
	GetOrdersProductsDetail(id int) (OrdersProducts, error)
}
