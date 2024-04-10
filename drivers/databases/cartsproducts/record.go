package cartsproducts

import "florist-gin/business/cartsproducts"

type CartsProducts struct {
	Id        int `gorm:"primaryKey"`
	CartId    int
	ProductId int
	Quantity  int
}

func (cartsProducts CartsProducts) ToUseCase() cartsproducts.CartsProducts {
	return cartsproducts.CartsProducts{
		Id:        cartsProducts.Id,
		CartId:    cartsProducts.CartId,
		ProductId: cartsProducts.ProductId,
		Quantity:  cartsProducts.Quantity,
	}
}

func FromUsecase(cartsProducts cartsproducts.CartsProducts) CartsProducts {
	return CartsProducts{
		Id:        cartsProducts.Id,
		CartId:    cartsProducts.CartId,
		ProductId: cartsProducts.ProductId,
		Quantity:  cartsProducts.Quantity,
	}
}
