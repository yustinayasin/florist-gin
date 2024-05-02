package ordersproducts

import (
	"errors"
	"florist-gin/business/ordersproducts"

	"gorm.io/gorm"
)

type OrdersProductsRepository struct {
	Db *gorm.DB
}

func NewOrdersProductsRepository(database *gorm.DB) ordersproducts.OrdersProductsRepoInterface {
	return &OrdersProductsRepository{
		Db: database,
	}
}

func (repo *OrdersProductsRepository) AddOrdersProducts(ordersProducts ordersproducts.OrdersProducts) (ordersproducts.OrdersProducts, error) {
	ordersProductsDB := FromUsecase(ordersProducts)

	resultOrdersProducts := repo.Db.Create(&ordersProductsDB)

	if resultOrdersProducts.Error != nil {
		return ordersproducts.OrdersProducts{}, resultOrdersProducts.Error
	}

	return ordersProductsDB.ToUsecase(), nil
}

func (repo *OrdersProductsRepository) EditOrdersProducts(ordersProducts ordersproducts.OrdersProducts, id int) (ordersproducts.OrdersProducts, error) {
	ordersProductsDb := FromUsecase(ordersProducts)

	var newOrdersProducts OrdersProducts

	resultOrdersProducts := repo.Db.First(&newOrdersProducts, id)

	if resultOrdersProducts.Error != nil {
		if resultOrdersProducts.Error == gorm.ErrRecordNotFound {
			return ordersproducts.OrdersProducts{}, errors.New("OrdersProducts not found")
		}
		return ordersproducts.OrdersProducts{}, errors.New("error in database")
	}

	newOrdersProducts.Quantity = ordersProductsDb.Quantity
	newOrdersProducts.Price = ordersProductsDb.Price

	repo.Db.Save(&newOrdersProducts)
	return newOrdersProducts.ToUsecase(), nil
}

func (repo *OrdersProductsRepository) DeleteOrdersProducts(id int) (ordersproducts.OrdersProducts, error) {
	var ordersProductsDb OrdersProducts

	resultFind := repo.Db.First(&ordersProductsDb, id)

	if resultFind.Error != nil {
		return ordersproducts.OrdersProducts{}, errors.New("ordersproducts not found")
	}

	result := repo.Db.Delete(&ordersProductsDb, id)

	if result.Error != nil {
		return ordersproducts.OrdersProducts{}, errors.New("ordersproducts not found")
	}

	return ordersProductsDb.ToUsecase(), nil
}

func (repo *OrdersProductsRepository) GetOrdersProductsDetail(id int) (ordersproducts.OrdersProducts, error) {
	var ordersProductsDb OrdersProducts

	resultFind := repo.Db.Preload("Order").Preload("Product").First(&ordersProductsDb, id)

	if resultFind.Error != nil {
		return ordersproducts.OrdersProducts{}, errors.New("order not found")
	}

	return ordersProductsDb.ToUsecase(), nil
}
