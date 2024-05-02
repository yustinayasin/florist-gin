package request

import (
	"florist-gin/business/orders"
	"strconv"
	"time"
)

type Order struct {
	Status     string    `form:"status"`
	Date       time.Time `form:"time"`
	TotalPrice int       `form:"totalPrice"`
	UserId     int       `form:"userId"`
}

func (order *Order) ToUsecase() *orders.Order {
	status, _ := strconv.ParseBool(order.Status)

	return &orders.Order{
		Status:     status,
		Date:       order.Date,
		TotalPrice: order.TotalPrice,
		UserId:     order.UserId,
	}
}
