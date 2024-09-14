package cartsproducts

import (
	"errors"
	"florist-gin/business/cartsproducts"
	"florist-gin/drivers/databases/carts"
	cartsProductsDB "florist-gin/drivers/databases/cartsproducts"
	productsDB "florist-gin/drivers/databases/products"
	cartRepo "florist-gin/drivers/repositories/carts"
	"florist-gin/drivers/repositories/products"

	"gorm.io/gorm"
)

type CartsProductsRepository struct {
	Db          *gorm.DB
	CartRepo    cartRepo.CartRepository
	ProductRepo products.ProductRepository
}

func NewCartsProductsRepository(database *gorm.DB, cartRepo cartRepo.CartRepository, productRepo products.ProductRepository) cartsproducts.CartsProductsRepoInterface {
	return &CartsProductsRepository{
		Db:          database,
		CartRepo:    cartRepo,
		ProductRepo: productRepo,
	}
}

func (repo *CartsProductsRepository) AddProductToCart(cartsProducts cartsproducts.CartsProducts) (cartsproducts.CartsProducts, error) {
	cartsProductsDB := cartsProductsDB.FromUsecase(cartsProducts)

	var cart carts.Cart
	var product productsDB.Product

	result := repo.CartRepo.Db.First(&cart, cartsProducts.CartId)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return cartsproducts.CartsProducts{}, errors.New("cart not found")
		}
		return cartsproducts.CartsProducts{}, errors.New("error in database")
	}

	result = repo.ProductRepo.Db.First(&product, cartsProducts.ProductId)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return cartsproducts.CartsProducts{}, errors.New("product not found")
		}
		return cartsproducts.CartsProducts{}, errors.New("error in database")
	}

	resultCartsProducts := repo.Db.Create(&cartsProductsDB)

	if resultCartsProducts.Error != nil {
		return cartsproducts.CartsProducts{}, result.Error
	}

	return cartsProductsDB.ToUseCase(), nil
}

func (repo *CartsProductsRepository) EditProductFromCart(cartsProducts cartsproducts.CartsProducts, id int) (cartsproducts.CartsProducts, error) {
	cartsProductsDb := cartsProductsDB.FromUsecase(cartsProducts)

	var cart carts.Cart
	var product productsDB.Product

	result := repo.CartRepo.Db.First(&cart, cartsProducts.CartId)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return cartsproducts.CartsProducts{}, errors.New("cart not found")
		}
		return cartsproducts.CartsProducts{}, errors.New("error in database")
	}

	result = repo.ProductRepo.Db.First(&product, cartsProducts.ProductId)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return cartsproducts.CartsProducts{}, errors.New("product not found")
		}
		return cartsproducts.CartsProducts{}, errors.New("error in database")
	}

	var newCartsProducts cartsProductsDB.CartsProducts

	resultCartsProducts := repo.Db.First(&newCartsProducts, id)

	if resultCartsProducts.Error != nil {
		if resultCartsProducts.Error == gorm.ErrRecordNotFound {
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
	var cartsProductsDb cartsProductsDB.CartsProducts

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
