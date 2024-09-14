package cartsproducts

import (
	"errors"
	"florist-gin/business/cartsproducts"
	"florist-gin/drivers/databases/carts"
	cPDB "florist-gin/drivers/databases/cartsproducts"
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

func (repo *CartsProductsRepository) AddProductToCart(cartsProducts cartsproducts.CartsProducts, userId int) (cartsproducts.CartsProducts, error) {
	cartsProductsDB := cPDB.FromUsecase(cartsProducts)

	var cart carts.Cart
	var product productsDB.Product
	var newCartsProducts cPDB.CartsProducts

	result := repo.CartRepo.Db.Where("user_id = ?", userId).First(&cart)

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

	result = repo.Db.Where("cart_id = ?", cart.Id).Where("product_id = ?", product.Id).First(&newCartsProducts)

	if result.Error == nil {
		newCartsProducts.Quantity = newCartsProducts.Quantity + 1
		editedCartsProducts, _ := repo.EditProductFromCart(newCartsProducts.ToUseCase(), newCartsProducts.Id)
		cartsProductsDB = cPDB.FromUsecase(editedCartsProducts)
	} else {
		cartsProductsDB.CartId = cart.Id
		resultCartsProducts := repo.Db.Create(&cartsProductsDB)

		if resultCartsProducts.Error != nil {
			return cartsproducts.CartsProducts{}, result.Error
		}
	}

	return cartsProductsDB.ToUseCase(), nil
}

func (repo *CartsProductsRepository) EditProductFromCart(cartsProducts cartsproducts.CartsProducts, id int) (cartsproducts.CartsProducts, error) {
	cartsProductsDb := cPDB.FromUsecase(cartsProducts)

	// var cart carts.Cart
	var product productsDB.Product

	// result := repo.CartRepo.Db.First(&cart, cartsProducts.CartId)

	// if result.Error != nil {
	// 	if result.Error == gorm.ErrRecordNotFound {
	// 		return cartsproducts.CartsProducts{}, errors.New("cart not found")
	// 	}
	// 	return cartsproducts.CartsProducts{}, errors.New("error in database")
	// }

	var newCartsProducts cPDB.CartsProducts

	resultCartsProducts := repo.Db.First(&newCartsProducts, id)

	if resultCartsProducts.Error != nil {
		if resultCartsProducts.Error == gorm.ErrRecordNotFound {
			return cartsproducts.CartsProducts{}, errors.New("CartsProducts not found")
		}
		return cartsproducts.CartsProducts{}, errors.New("error in database")
	}

	result := repo.ProductRepo.Db.First(&product, newCartsProducts.ProductId)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return cartsproducts.CartsProducts{}, errors.New("product not found")
		}
		return cartsproducts.CartsProducts{}, errors.New("error in database")
	}

	if cartsProductsDb.Quantity > product.Stock {
		return cartsproducts.CartsProducts{}, errors.New("the requested quantity exceeds available stock")
	}

	// newCartsProducts.CartId = cartsProductsDb.CartId
	// newCartsProducts.ProductId = cartsProductsDb.ProductId
	newCartsProducts.Quantity = cartsProductsDb.Quantity

	repo.Db.Save(&newCartsProducts)
	return newCartsProducts.ToUseCase(), nil
}

func (repo *CartsProductsRepository) DeleteProductFromCart(id int) (cartsproducts.CartsProducts, error) {
	var cartsProductsDb cPDB.CartsProducts

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
