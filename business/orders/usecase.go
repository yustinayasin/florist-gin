package orders

import (
	"errors"
	"time"
)

type OrderUseCase struct {
	Repo OrderRepoInterface
}

func NewUseCase(orderRepo OrderRepoInterface) OrderUseCaseInterface {
	return &OrderUseCase{
		Repo: orderRepo,
	}
}

func (orderUseCase *OrderUseCase) AddOrder(order Order) (Order, error) {
	if order.TotalPrice == 0 {
		return Order{}, errors.New("total price cannot be empty")
	}

	if order.UserId == 0 {
		return Order{}, errors.New("user ID cannot be empty")
	}

	order.Status = false
	order.Date = time.Now()

	orderRepo, err := orderUseCase.Repo.AddOrder(order)

	if err != nil {
		return Order{}, err
	}

	return orderRepo, nil
}

func (orderUseCase *OrderUseCase) EditOrder(order Order, id int) (Order, error) {
	if id == 0 {
		return Order{}, errors.New("order ID cannot be empty")
	}

	orderRepo, err := orderUseCase.Repo.EditOrder(order, id)

	if err != nil {
		return Order{}, err
	}

	return orderRepo, nil
}

func (orderUseCase *OrderUseCase) DeleteOrder(id int) (Order, error) {
	if id == 0 {
		return Order{}, errors.New("order ID cannot be empty")
	}

	orderRepo, err := orderUseCase.Repo.DeleteOrder(id)

	if err != nil {
		return Order{}, err
	}

	return orderRepo, nil
}

func (orderUseCase *OrderUseCase) GetOrderDetail(id int) (Order, error) {
	if id == 0 {
		return Order{}, errors.New("order ID cannot be empty")
	}

	orderRepo, err := orderUseCase.Repo.GetOrderDetail(id)

	if err != nil {
		return Order{}, err
	}

	return orderRepo, nil
}
