package carts

import (
	"florist-gin/business/carts"
	"florist-gin/drivers/databases/products"
)

type Cart struct {
	Id      int                `gorm:"primaryKey;unique;autoIncrement:true"`
	UserId  int                `gorm:"unique"`
	Product []products.Product `gorm:"many2many:CartsProducts;"`
}

func (cart Cart) ToUseCase() carts.Cart {
	newProducts := products.ToUsecaseList(cart.Product)

	return carts.Cart{
		Id:      cart.Id,
		UserId:  cart.UserId,
		Product: newProducts,
	}
}

func FromUsecase(cart carts.Cart) Cart {
	newProducts := products.FromUsecaseList(cart.Product)

	return Cart{
		Id:      cart.Id,
		UserId:  cart.UserId,
		Product: newProducts,
	}
}
