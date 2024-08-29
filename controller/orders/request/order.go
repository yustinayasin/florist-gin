package request

import (
	"florist-gin/business/orders"
	"time"
)

type Order struct {
	Date       time.Time `form:"time"`
	TotalPrice int       `form:"totalPrice"`
	UserId     int       `form:"userId"`
}

func (order *Order) ToUsecase() *orders.Order {
	return &orders.Order{
		Date:       order.Date,
		TotalPrice: order.TotalPrice,
		UserId:     order.UserId,
	}
}
