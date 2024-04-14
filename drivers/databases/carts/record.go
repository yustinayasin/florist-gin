package carts

import (
	"florist-gin/business/carts"
	"florist-gin/drivers/databases/products"
	"time"
)

type Cart struct {
	Id        uint32             `gorm:"primaryKey;unique"`
	UserId    uint32             `gorm:"unique"`
	Product   []products.Product `gorm:"many2many:CartsProducts;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (cart Cart) ToUseCase() carts.Cart {
	newProducts := products.ToUsecaseList(cart.Product)

	return carts.Cart{
		Id:        cart.Id,
		UserId:    cart.UserId,
		Product:   newProducts,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
	}
}

func FromUsecase(cart carts.Cart) Cart {
	newProducts := products.FromUsecaseList(cart.Product)

	return Cart{
		Id:        cart.Id,
		UserId:    cart.UserId,
		Product:   newProducts,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
	}
}
