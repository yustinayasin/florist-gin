package response

import (
	"florist-gin/business/orders"
	"florist-gin/business/users"
	"time"
)

type OrderResponse struct {
	Id         uint32     `json:"id"`
	Status     bool       `json:"status"`
	Date       time.Time  `json:"time"`
	TotalPrice int        `json:"totalPrice"`
	UserId     uint32     `json:"userId"`
	User       users.User `json:"user"`
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
