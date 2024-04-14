package products

import (
	"florist-gin/business/categories"
	"mime/multipart"
	"time"
)

type Product struct {
	Id          uint32
	Name        string
	Description string
	Price       int
	Stock       int
	FileName    string
	File        multipart.File
	CategoryId  uint32
	Category    categories.Category
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductUseCaseInterface interface {
	AddProduct(product Product) (Product, error)
	EditProduct(product Product, id uint32) (Product, error)
	DeleteProduct(id uint32) (Product, error)
	GetProductDetail(id uint32) (Product, error)
}

type ProductRepoInterface interface {
	AddProduct(product Product) (Product, error)
	EditProduct(product Product, id uint32) (Product, error)
	DeleteProduct(id uint32) (Product, error)
	GetProductDetail(id uint32) (Product, error)
}
