package response

import (
	"florist-gin/business/categories"
	"florist-gin/business/products"
)

type ProductResponse struct {
	Id          int                 `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Price       int                 `json:"price"`
	Stock       int                 `json:"stock"`
	CategoryId  int                 `json:"categoryId"`
	Category    categories.Category `json:"category"`
}

func FromUsecase(product products.Product) ProductResponse {
	return ProductResponse{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		CategoryId:  product.CategoryId,
	}
}

func FromUsecaseList(product []products.Product) []ProductResponse {
	var productResponse []ProductResponse

	for _, v := range product {
		productResponse = append(productResponse, FromUsecase(v))
	}

	return productResponse
}
