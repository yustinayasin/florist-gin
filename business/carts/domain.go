package carts

import (
	"florist-gin/business/products"
	"time"
)

type Cart struct {
	Id        uint32
	UserId    uint32
	Product   []products.Product
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartUseCaseInterface interface {
	GetCart(id uint32) (Cart, error)
}

type CartRepoInterface interface {
	GetCart(id uint32) (Cart, error)
}
