package cartsproducts

import (
	"errors"
	"florist-gin/business/cartsproducts"

	"gorm.io/gorm"
)

type CartsProductsRepository struct {
	Db *gorm.DB
}

func NewCartsProductsRepository(database *gorm.DB) cartsproducts.CartsProductsRepoInterface {
	return &CartsProductsRepository{
		Db: database,
	}
}

func (repo *CartsProductsRepository) AddProductToCart(cartsProducts cartsproducts.CartsProducts) (cartsproducts.CartsProducts, error) {
	cartsProductsDB := FromUsecase(cartsProducts)

	result := repo.Db.Create(&cartsProductsDB)

	if result.Error != nil {
		return cartsproducts.CartsProducts{}, result.Error
	}

	return cartsProductsDB.ToUseCase(), nil
}

func (repo *CartsProductsRepository) EditProductFromCart(cartsProducts cartsproducts.CartsProducts, id int) (cartsproducts.CartsProducts, error) {
	cartsProductsDb := FromUsecase(cartsProducts)

	var newCartsProducts CartsProducts

	result := repo.Db.First(&newCartsProducts, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return cartsproducts.CartsProducts{}, errors.New("CartsProducts not found")
		}
		return cartsproducts.CartsProducts{}, errors.New("error in database")
	}

	newCartsProducts.CartId = cartsProductsDb.CartId
	newCartsProducts.ProductId = cartsProductsDb.ProductId
	newCartsProducts.Quantity = cartsProductsDb.Quantity

	repo.Db.Save(&newCartsProducts)
	return newCartsProducts.ToUseCase(), nil
}

func (repo *CartsProductsRepository) DeleteProductFromCart(id int) (cartsproducts.CartsProducts, error) {
	var cartsProductsDb CartsProducts

	resultFind := repo.Db.First(&cartsProductsDb, id)

	if resultFind.Error != nil {
		return cartsproducts.CartsProducts{}, errors.New("cartsproducts not found")
	}

	result := repo.Db.Delete(&cartsProductsDb, id)

	if result.Error != nil {
		return cartsproducts.CartsProducts{}, errors.New("cartsproducts not found")
	}

	return cartsProductsDb.ToUseCase(), nil
}
