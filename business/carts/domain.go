package carts

import (
	"time"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       int
	Stock       int
	FileUrl     string
	Quantity    int
}

type Cart struct {
	Id         int
	UserId     int
	Products   []Product
	TotalPrice int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CartUseCaseInterface interface {
	GetCart(userId int) (Cart, error)
}

type CartRepoInterface interface {
	GetCart(userId int) (Cart, error)
}
