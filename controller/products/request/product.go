package request

import (
	"florist-gin/business/products"
)

type Product struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	Price       int    `form:"price"`
	Stock       int    `form:"stock"`
	FileName    string
	File        []byte
	CategoryId  int `form:"categoryId"`
}

func (product *Product) ToUsecase() *products.Product {
	return &products.Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		FileName:    product.FileName,
		File:        product.File,
		CategoryId:  product.CategoryId,
	}
}
