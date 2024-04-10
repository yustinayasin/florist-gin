package products

import "florist-gin/business/categories"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       int
	Stock       int
	CategoryId  int
	Category    categories.Category
}

type ProductUseCaseInterface interface {
	AddProduct(product Product) (Product, error)
	EditProduct(product Product, id int) (Product, error)
	DeleteProduct(id int) (Product, error)
	GetProductDetail(id int) (Product, error)
}

type ProductRepoInterface interface {
	AddProduct(product Product) (Product, error)
	EditProduct(product Product, id int) (Product, error)
	DeleteProduct(id int) (Product, error)
	GetProductDetail(id int) (Product, error)
}
