package response

import (
	"florist-gin/business/categories"
	"florist-gin/business/products"
	"time"
)

type ProductResponse struct {
	Id          int    `form:"id"`
	Name        string `form:"name"`
	Description string `form:"description"`
	Price       int    `form:"price"`
	Stock       int    `form:"stock"`
	FileName    string `form:"fileName"`
	FileUrl     string
	CategoryId  int                 `form:"categoryId"`
	Category    categories.Category `form:"category"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func FromUsecase(product products.Product) ProductResponse {
	return ProductResponse{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		FileName:    product.FileName,
		FileUrl:     product.FileUrl.String(),
		CategoryId:  product.CategoryId,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func FromUsecaseList(product []products.Product) []ProductResponse {
	var productResponse []ProductResponse

	for _, v := range product {
		productResponse = append(productResponse, FromUsecase(v))
	}

	return productResponse
}
