package response

import (
	"florist-gin/business/categories"
	"florist-gin/business/products"
	"time"
)

type ProductResponse struct {
	Id          uint32              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Price       int                 `json:"price"`
	Stock       int                 `json:"stock"`
	FileName    string              `json:"fileName"`
	CategoryId  uint32              `json:"categoryId"`
	Category    categories.Category `json:"category"`
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
