package request

import (
	"florist-gin/business/orders"
	"time"
)

type Order struct {
	Status     bool      `json:"status"`
	Date       time.Time `json:"time"`
	TotalPrice int       `json:"totalPrice"`
	UserId     uint32    `json:"userId"`
}

func (order *Order) ToUsecase() *orders.Order {
	return &orders.Order{
		Status:     order.Status,
		Date:       order.Date,
		TotalPrice: order.TotalPrice,
		UserId:     order.UserId,
	}
}
