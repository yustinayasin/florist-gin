package products

import (
	"florist-gin/business/products"
	"florist-gin/drivers/databases/categories"
	"time"
)

type Product struct {
	Id          int `gorm:"primaryKey;unique"`
	Name        string
	Description string
	Price       int
	Stock       int
	FileName    string
	CategoryId  int
	Category    categories.Category `gorm:"foreignKey:CategoryId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (product Product) ToUsecase() products.Product {
	newCategory := product.Category.ToUseCase()

	return products.Product{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		FileName:    product.FileName,
		CategoryId:  product.CategoryId,
		Category:    newCategory,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func ToUsecaseList(product []Product) []products.Product {
	var newProducts []products.Product

	for _, v := range product {
		newProducts = append(newProducts, v.ToUsecase())
	}
	return newProducts
}

func FromUsecase(product products.Product) Product {
	newCategory := categories.FromUsecase(product.Category)

	return Product{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		FileName:    product.FileName,
		CategoryId:  product.CategoryId,
		Category:    newCategory,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func FromUsecaseList(products []products.Product) []Product {
	var newProducts []Product

	for _, v := range products {
		newProducts = append(newProducts, FromUsecase(v))
	}

	return newProducts
}
