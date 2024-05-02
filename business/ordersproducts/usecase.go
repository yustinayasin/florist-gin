package ordersproducts

import (
	"errors"
)

type OrdersProductsUseCase struct {
	Repo OrdersProductsRepoInterface
}

func NewUseCase(ordersProductsRepo OrdersProductsRepoInterface) OrdersProductsUseCaseInterface {
	return &OrdersProductsUseCase{
		Repo: ordersProductsRepo,
	}
}

func (ordersProductsUseCase *OrdersProductsUseCase) AddOrdersProducts(ordersproducts OrdersProducts) (OrdersProducts, error) {
	if ordersproducts.OrderId == 0 {
		return OrdersProducts{}, errors.New("order id cannot be empty")
	}

	if ordersproducts.ProductId == 0 {
		return OrdersProducts{}, errors.New("product id cannot be empty")
	}

	if ordersproducts.Quantity == 0 {
		return OrdersProducts{}, errors.New("quantity cannot be empty")
	}

	if ordersproducts.Price == 0 {
		return OrdersProducts{}, errors.New("price cannot be empty")
	}

	userRepo, err := ordersProductsUseCase.Repo.AddOrdersProducts(ordersproducts)

	if err != nil {
		return OrdersProducts{}, err
	}

	return userRepo, nil
}

func (ordersProductsUseCase *OrdersProductsUseCase) EditOrdersProducts(ordersproducts OrdersProducts, id int) (OrdersProducts, error) {
	if id == 0 {
		return OrdersProducts{}, errors.New("orders products ID cannot be empty")
	}

	if ordersproducts.Quantity == 0 {
		return OrdersProducts{}, errors.New("quantity cannot be empty")
	}

	if ordersproducts.Price == 0 {
		return OrdersProducts{}, errors.New("price cannot be empty")
	}

	ordersProductsRepo, err := ordersProductsUseCase.Repo.EditOrdersProducts(ordersproducts, id)

	if err != nil {
		return OrdersProducts{}, err
	}

	return ordersProductsRepo, nil
}

func (ordersProductsUseCase *OrdersProductsUseCase) DeleteOrdersProducts(id int) (OrdersProducts, error) {
	if id == 0 {
		return OrdersProducts{}, errors.New("orders products ID cannot be empty")
	}

	ordersProductsRepo, err := ordersProductsUseCase.Repo.DeleteOrdersProducts(id)

	if err != nil {
		return OrdersProducts{}, err
	}

	return ordersProductsRepo, nil
}

func (ordersProductsUseCase *OrdersProductsUseCase) GetOrdersProductsDetail(id int) (OrdersProducts, error) {
	if id == 0 {
		return OrdersProducts{}, errors.New("orders products ID cannot be empty")
	}

	ordersProductsRepo, err := ordersProductsUseCase.Repo.GetOrdersProductsDetail(id)

	if err != nil {
		return OrdersProducts{}, err
	}

	return ordersProductsRepo, nil
}
