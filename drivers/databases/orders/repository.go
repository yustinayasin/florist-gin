package orders

import (
	"errors"
	"florist-gin/business/orders"

	"gorm.io/gorm"
)

type OrderRepository struct {
	Db *gorm.DB
}

func NewOrderRepository(database *gorm.DB) orders.OrderRepoInterface {
	return &OrderRepository{
		Db: database,
	}
}

func (repo *OrderRepository) AddOrder(order orders.Order) (orders.Order, error) {
	orderDB := FromUsecase(order)

	result := repo.Db.Create(&orderDB)

	if result.Error != nil {
		return orders.Order{}, result.Error
	}

	return orderDB.ToUsecase(), nil
}

func (repo *OrderRepository) EditOrder(order orders.Order, id int) (orders.Order, error) {
	orderDb := FromUsecase(order)

	var newOrder Order

	result := repo.Db.Preload("User").First(&newOrder, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return orders.Order{}, errors.New("Order not found")
		}
		return orders.Order{}, errors.New("error in database")
	}

	newOrder.TotalPrice = orderDb.TotalPrice
	repo.Db.Save(&newOrder)
	return newOrder.ToUsecase(), nil
}

func (repo *OrderRepository) DeleteOrder(id int) (orders.Order, error) {
	var orderDb Order

	resultFind := repo.Db.Preload("User").First(&orderDb, id)

	if resultFind.Error != nil {
		return orders.Order{}, errors.New("order not found")
	}

	result := repo.Db.Delete(&orderDb, id)

	if result.Error != nil {
		return orders.Order{}, errors.New("order not found")
	}

	return orderDb.ToUsecase(), nil
}

func (repo *OrderRepository) GetOrderDetail(id int) (orders.Order, error) {
	var orderDb Order

	resultFind := repo.Db.Preload("User").First(&orderDb, id)

	if resultFind.Error != nil {
		return orders.Order{}, errors.New("order not found")
	}

	return orderDb.ToUsecase(), nil
}
