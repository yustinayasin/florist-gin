package products

import (
	"florist-gin/business/categories"
	"net/url"
	"time"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       int
	Stock       int
	FileName    string
	File        []byte
	FileUrl     *url.URL
	CategoryId  int
	Category    categories.Category
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductUseCaseInterface interface {
	AddProduct(product Product) (Product, error)
	EditProduct(product Product, id int) (Product, error)
	DeleteProduct(id int) (Product, error)
	GetProductDetail(id int) (Product, error)
	GetAllProduct(categoryId int) ([]Product, error)
}

type ProductRepoInterface interface {
	AddProduct(product Product) (Product, error)
	EditProduct(product Product, id int) (Product, error)
	DeleteProduct(id int) (Product, error)
	GetProductDetail(id int) (Product, error)
	GetAllProduct(categoryId int) ([]Product, error)
}
