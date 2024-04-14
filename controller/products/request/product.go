package request

import (
	"florist-gin/business/products"
	"mime/multipart"
)

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	FileName    string
	File        multipart.File
	CategoryId  uint32 `json:"categoryId"`
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
