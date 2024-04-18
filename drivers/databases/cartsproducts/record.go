package cartsproducts

import (
	"florist-gin/business/cartsproducts"
	"time"
)

type CartsProducts struct {
	Id        int `gorm:"primaryKey"`
	CartId    int
	ProductId int
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (cartsProducts CartsProducts) ToUseCase() cartsproducts.CartsProducts {
	return cartsproducts.CartsProducts{
		Id:        cartsProducts.Id,
		CartId:    cartsProducts.CartId,
		ProductId: cartsProducts.ProductId,
		Quantity:  cartsProducts.Quantity,
		CreatedAt: cartsProducts.CreatedAt,
		UpdatedAt: cartsProducts.UpdatedAt,
	}
}

func FromUsecase(cartsProducts cartsproducts.CartsProducts) CartsProducts {
	return CartsProducts{
		Id:        cartsProducts.Id,
		CartId:    cartsProducts.CartId,
		ProductId: cartsProducts.ProductId,
		Quantity:  cartsProducts.Quantity,
		CreatedAt: cartsProducts.CreatedAt,
		UpdatedAt: cartsProducts.UpdatedAt,
	}
}
