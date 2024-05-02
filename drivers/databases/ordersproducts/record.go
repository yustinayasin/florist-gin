package ordersproducts

import (
	"florist-gin/business/ordersproducts"
	"florist-gin/drivers/databases/orders"
	"florist-gin/drivers/databases/products"
	"time"
)

type OrdersProducts struct {
	Id        int `gorm:"primaryKey;unique"`
	Quantity  int
	Price     int
	OrderId   int
	ProductId int
	Order     orders.Order     `gorm:"foreignKey:OrderId"`
	Product   products.Product `gorm:"foreignKey:ProductId"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (orderproduct OrdersProducts) ToUsecase() ordersproducts.OrdersProducts {
	newOrder := orders.Order.ToUsecase(orderproduct.Order)
	newProduct := products.Product.ToUsecase(orderproduct.Product)

	return ordersproducts.OrdersProducts{
		Id:        orderproduct.Id,
		Quantity:  orderproduct.Quantity,
		Price:     orderproduct.Price,
		OrderId:   orderproduct.OrderId,
		ProductId: orderproduct.ProductId,
		Order:     newOrder,
		Product:   newProduct,
		CreatedAt: orderproduct.CreatedAt,
		UpdatedAt: orderproduct.UpdatedAt,
	}
}

func ToUsecaseList(orderproduct []OrdersProducts) []ordersproducts.OrdersProducts {
	var newOrders []ordersproducts.OrdersProducts

	for _, v := range orderproduct {
		newOrders = append(newOrders, v.ToUsecase())
	}
	return newOrders
}

func FromUsecase(orderproduct ordersproducts.OrdersProducts) OrdersProducts {
	newOrder := orders.FromUsecase(orderproduct.Order)
	newProduct := products.FromUsecase(orderproduct.Product)

	return OrdersProducts{
		Id:        orderproduct.Id,
		Quantity:  orderproduct.Quantity,
		Price:     orderproduct.Price,
		OrderId:   orderproduct.OrderId,
		ProductId: orderproduct.ProductId,
		Order:     newOrder,
		Product:   newProduct,
		CreatedAt: orderproduct.CreatedAt,
		UpdatedAt: orderproduct.UpdatedAt,
	}
}
