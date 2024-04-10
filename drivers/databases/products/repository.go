package products

import (
	"errors"
	"florist-gin/business/products"

	"gorm.io/gorm"
)

type ProductRepository struct {
	Db *gorm.DB
}

func NewProductRepository(database *gorm.DB) products.ProductRepoInterface {
	return &ProductRepository{
		Db: database,
	}
}

func (repo *ProductRepository) AddProduct(product products.Product) (products.Product, error) {
	productDB := FromUsecase(product)

	result := repo.Db.Create(&productDB)

	if result.Error != nil {
		return products.Product{}, result.Error
	}

	return productDB.ToUsecase(), nil
}

func (repo *ProductRepository) EditProduct(product products.Product, id int) (products.Product, error) {
	productDb := FromUsecase(product)

	var newProduct Product

	result := repo.Db.First(&newProduct, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return products.Product{}, errors.New("Product not found")
		}
		return products.Product{}, errors.New("error in database")
	}

	newProduct.Name = productDb.Name
	newProduct.Description = productDb.Description
	newProduct.Price = productDb.Price
	newProduct.Stock = productDb.Stock
	newProduct.CategoryId = productDb.CategoryId

	repo.Db.Save(&newProduct)
	return newProduct.ToUsecase(), nil
}

func (repo *ProductRepository) DeleteProduct(id int) (products.Product, error) {
	var productDb Product

	resultFind := repo.Db.First(&productDb, id)

	if resultFind.Error != nil {
		return products.Product{}, errors.New("Product not found")
	}

	result := repo.Db.Delete(&productDb, id)

	if result.Error != nil {
		return products.Product{}, errors.New("Product not found")
	}

	return productDb.ToUsecase(), nil
}

func (repo *ProductRepository) GetProductDetail(id int) (products.Product, error) {
	var productDb Product

	resultFind := repo.Db.Preload("Category").First(&productDb, id)

	if resultFind.Error != nil {
		return products.Product{}, errors.New("Product not found")
	}

	return productDb.ToUsecase(), nil
}
