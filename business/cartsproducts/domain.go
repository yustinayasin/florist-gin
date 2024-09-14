package cartsproducts

import (
	"florist-gin/business/products"
	"time"
)

type CartsProducts struct {
	Id        int
	CartId    int
	ProductId int
	Quantity  int
	Product   products.Product
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartsProductsUseCaseInterface interface {
	AddProductToCart(cartsProducts CartsProducts, userId int) (CartsProducts, error)
	EditProductFromCart(cartsProducts CartsProducts, idCartsProducts int) (CartsProducts, error)
	DeleteProductFromCart(idCartsProducts int) (CartsProducts, error)
}

type CartsProductsRepoInterface interface {
	AddProductToCart(cartsProducts CartsProducts, userId int) (CartsProducts, error)
	EditProductFromCart(cartsProducts CartsProducts, idCartsProducts int) (CartsProducts, error)
	DeleteProductFromCart(idCartsProducts int) (CartsProducts, error)
}
