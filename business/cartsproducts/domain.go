package cartsproducts

type CartsProducts struct {
	Id        int
	CartId    int
	ProductId int
	Quantity  int
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
