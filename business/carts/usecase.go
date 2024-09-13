package carts

import (
	"errors"
)

type CartUseCase struct {
	Repo CartRepoInterface
}

func NewUseCase(cartRepo CartRepoInterface) CartUseCaseInterface {
	return &CartUseCase{
		Repo: cartRepo,
	}
}

func (cartUseCase *CartUseCase) GetCart(userId int) (Cart, error) {
	if userId == 0 {
		return Cart{}, errors.New("cart user ID cannot be empty")
	}

	cartRepo, err := cartUseCase.Repo.GetCart(userId)

	if err != nil {
		return Cart{}, err
	}

	return cartRepo, nil
}
