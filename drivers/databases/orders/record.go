package orders

import (
	"florist-gin/business/orders"
	"florist-gin/drivers/databases/users"
	"time"
)

type Order struct {
	Id         int `gorm:"primaryKey;unique"`
	Date       time.Time
	TotalPrice int
	UserId     int
	User       users.User `gorm:"foreignKey:UserId"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (order Order) ToUsecase() orders.Order {
	newUser := users.User.ToUsecase(order.User)

	return orders.Order{
		Id:         order.Id,
		Date:       order.Date,
		TotalPrice: order.TotalPrice,
		UserId:     order.UserId,
		User:       newUser,
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
	}
}

func ToUsecaseList(order []Order) []orders.Order {
	var newOrders []orders.Order

	for _, v := range order {
		newOrders = append(newOrders, v.ToUsecase())
	}
	return newOrders
}

func FromUsecase(order orders.Order) Order {
	newUser := users.FromUsecase(order.User)

	return Order{
		Id:         order.Id,
		Date:       order.Date,
		TotalPrice: order.TotalPrice,
		UserId:     order.UserId,
		User:       newUser,
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
	}
}
