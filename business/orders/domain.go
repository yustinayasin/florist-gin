package orders

import (
	"florist-gin/business/users"
	"time"
)

type Order struct {
	Id         uint32
	Status     bool
	Date       time.Time
	TotalPrice int
	UserId     uint32
	User       users.User
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type OrderUseCaseInterface interface {
	AddOrder(order Order) (Order, error)
	EditOrder(order Order, id uint32) (Order, error)
	DeleteOrder(id uint32) (Order, error)
	GetOrderDetail(id uint32) (Order, error)
}

type OrderRepoInterface interface {
	AddOrder(order Order) (Order, error)
	EditOrder(order Order, id uint32) (Order, error)
	DeleteOrder(id uint32) (Order, error)
	GetOrderDetail(id uint32) (Order, error)
}
