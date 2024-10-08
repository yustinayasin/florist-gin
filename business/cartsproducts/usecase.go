package cartsproducts

import (
	"errors"
)

type CartsProductsUseCase struct {
	Repo CartsProductsRepoInterface
}

func NewUseCase(cartsProductsRepo CartsProductsRepoInterface) CartsProductsUseCaseInterface {
	return &CartsProductsUseCase{
		Repo: cartsProductsRepo,
	}
}

func (cartsProductsUseCase *CartsProductsUseCase) AddProductToCart(cartsproducts CartsProducts, userId int) (CartsProducts, error) {
	if userId == 0 {
		return CartsProducts{}, errors.New("user id cannot be empty")
	}

	if cartsproducts.ProductId == 0 {
		return CartsProducts{}, errors.New("product id cannot be empty")
	}

	if cartsproducts.Quantity == 0 {
		return CartsProducts{}, errors.New("quantity cannot be empty")
	}

	userRepo, err := cartsProductsUseCase.Repo.AddProductToCart(cartsproducts, userId)

	if err != nil {
		return CartsProducts{}, err
	}

	return userRepo, nil
}

func (cartsProductsUseCase *CartsProductsUseCase) EditProductFromCart(cartsproducts CartsProducts, id int) (CartsProducts, error) {
	if id == 0 {
		return CartsProducts{}, errors.New("carts products ID cannot be empty")
	}

	// if cartsproducts.CartId == 0 {
	// 	return CartsProducts{}, errors.New("card ID cannot be empty")
	// }

	// if cartsproducts.ProductId == 0 {
	// 	return CartsProducts{}, errors.New("product ID cannot be empty")
	// }

	if cartsproducts.Quantity == 0 {
		return CartsProducts{}, errors.New("quantity cannot be empty")
	}

	cartsProductsRepo, err := cartsProductsUseCase.Repo.EditProductFromCart(cartsproducts, id)

	if err != nil {
		return CartsProducts{}, err
	}

	return cartsProductsRepo, nil
}

func (cartsProductsUseCase *CartsProductsUseCase) DeleteProductFromCart(id int) (CartsProducts, error) {
	if id == 0 {
		return CartsProducts{}, errors.New("carts products ID cannot be empty")
	}

	cartsProductsRepo, err := cartsProductsUseCase.Repo.DeleteProductFromCart(id)

	if err != nil {
		return CartsProducts{}, err
	}

	return cartsProductsRepo, nil
}
