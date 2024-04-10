package request

import "florist-gin/business/products"

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	CategoryId  int    `json:"categoryId"`
}

func (product *Product) ToUsecase() *products.Product {
	return &products.Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		CategoryId:  product.CategoryId,
	}
}
