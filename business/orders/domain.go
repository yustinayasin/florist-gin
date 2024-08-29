package orders

import (
	"florist-gin/business/users"
	"time"
)

type Order struct {
	Id         int
	Date       time.Time
	TotalPrice int
	UserId     int
	User       users.User
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type OrderUseCaseInterface interface {
	AddOrder(order Order) (Order, error)
	EditOrder(order Order, id int) (Order, error)
	DeleteOrder(id int) (Order, error)
	GetOrderDetail(id int) (Order, error)
}

type OrderRepoInterface interface {
	AddOrder(order Order) (Order, error)
	EditOrder(order Order, id int) (Order, error)
	DeleteOrder(id int) (Order, error)
	GetOrderDetail(id int) (Order, error)
}
