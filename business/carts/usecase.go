package carts

import "errors"

type CartUseCase struct {
	Repo CartRepoInterface
}

func NewUseCase(cartRepo CartRepoInterface) CartUseCaseInterface {
	return &CartUseCase{
		Repo: cartRepo,
	}
}

func (cartUseCase *CartUseCase) GetCart(id int) (Cart, error) {
	if id == 0 {
		return Cart{}, errors.New("cart ID cannot be empty")
	}

	cartRepo, err := cartUseCase.Repo.GetCart(id)

	if err != nil {
		return Cart{}, err
	}

	return cartRepo, nil
}
