package cartsproducts

import (
	"florist-gin/business/cartsproducts"
	"florist-gin/drivers/databases/products"
	"time"
)

type CartsProducts struct {
	Id        int `gorm:"primaryKey"`
	CartId    int
	ProductId int
	Quantity  int
	Product   products.Product `gorm:"foreignKey:ProductId"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (cartsProducts CartsProducts) ToUseCase() cartsproducts.CartsProducts {
	return cartsproducts.CartsProducts{
		Id:        cartsProducts.Id,
		CartId:    cartsProducts.CartId,
		ProductId: cartsProducts.ProductId,
		Product:   cartsProducts.Product.ToUsecase(),
		Quantity:  cartsProducts.Quantity,
		CreatedAt: cartsProducts.CreatedAt,
		UpdatedAt: cartsProducts.UpdatedAt,
	}
}

// func ToUsecaseList(cartsProducts []CartsProducts) []carts.CartProduct {
// 	var newCartsProducts carts.CartProduct
// 	var newCartsProductsList []carts.CartProduct

// 	for _, v := range cartsProducts {
// 		newCartsProducts.Product = v.ToUseCase().Product
// 		newCartsProducts.Quantity = v.ToUseCase().Quantity
// 		newCartsProductsList = append(newCartsProductsList, newCartsProducts)
// 	}
// 	return newCartsProductsList
// }

func FromUsecase(cartsProducts cartsproducts.CartsProducts) CartsProducts {
	return CartsProducts{
		Id:        cartsProducts.Id,
		CartId:    cartsProducts.CartId,
		ProductId: cartsProducts.ProductId,
		Product:   products.FromUsecase(cartsProducts.Product),
		Quantity:  cartsProducts.Quantity,
		CreatedAt: cartsProducts.CreatedAt,
		UpdatedAt: cartsProducts.UpdatedAt,
	}
}
