package carts

import "florist-gin/business/products"

type Cart struct {
	Id      int
	UserId  int
	Product []products.Product
}

type CartUseCaseInterface interface {
	GetCart(id int) (Cart, error)
}

type CartRepoInterface interface {
	GetCart(id int) (Cart, error)
}
