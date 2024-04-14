package cartsproducts

import "time"

type CartsProducts struct {
	Id        uint32
	CartId    uint32
	ProductId uint32
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartsProductsUseCaseInterface interface {
	AddProductToCart(cartsProducts CartsProducts) (CartsProducts, error)
	EditProductFromCart(cartsProducts CartsProducts, idCartsProducts uint32) (CartsProducts, error)
	DeleteProductFromCart(idCartsProducts uint32) (CartsProducts, error)
}

type CartsProductsRepoInterface interface {
	AddProductToCart(cartsProducts CartsProducts) (CartsProducts, error)
	EditProductFromCart(cartsProducts CartsProducts, idCartsProducts uint32) (CartsProducts, error)
	DeleteProductFromCart(idCartsProducts uint32) (CartsProducts, error)
}
