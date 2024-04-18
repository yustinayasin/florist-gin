package response

import (
	"florist-gin/business/orders"
	"florist-gin/business/users"
	"time"
)

type OrderResponse struct {
	Id         int        `form:"id"`
	Status     bool       `form:"status"`
	Date       time.Time  `form:"time"`
	TotalPrice int        `form:"totalPrice"`
	UserId     int        `form:"userId"`
	User       users.User `form:"user"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func FromUsecase(order orders.Order) OrderResponse {
	return OrderResponse{
		Id:         order.Id,
		Status:     order.Status,
		Date:       order.Date,
		TotalPrice: order.TotalPrice,
		UserId:     order.UserId,
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
	}
}

func FromUsecaseList(order []orders.Order) []OrderResponse {
	var orderResponse []OrderResponse

	for _, v := range order {
		orderResponse = append(orderResponse, FromUsecase(v))
	}

	return orderResponse
}
