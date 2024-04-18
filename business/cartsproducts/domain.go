package cartsproducts

import "time"

type CartsProducts struct {
	Id        int
	CartId    int
	ProductId int
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartsProductsUseCaseInterface interface {
	AddProductToCart(cartsProducts CartsProducts) (CartsProducts, error)
	EditProductFromCart(cartsProducts CartsProducts, idCartsProducts int) (CartsProducts, error)
	DeleteProductFromCart(idCartsProducts int) (CartsProducts, error)
}

type CartsProductsRepoInterface interface {
	AddProductToCart(cartsProducts CartsProducts) (CartsProducts, error)
	EditProductFromCart(cartsProducts CartsProducts, idCartsProducts int) (CartsProducts, error)
	DeleteProductFromCart(idCartsProducts int) (CartsProducts, error)
}
