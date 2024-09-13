package carts

import (
	"florist-gin/business/products"
	"time"
)

type Cart struct {
	Id        int
	UserId    int
	Product   []products.Product
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartUseCaseInterface interface {
	GetCart(userId int) (Cart, error)
}

type CartRepoInterface interface {
	GetCart(userId int) (Cart, error)
}
