package response

import (
	"florist-gin/business/carts"
	"time"
)

type CartResponse struct {
	Id         int `form:"id"`
	UserId     int `form:"user_id"`
	Products   []carts.Product
	TotalPrice int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func FromUsecase(cart carts.Cart) CartResponse {
	// var newProducts CartProduct
	// var newProductsList []CartProduct
	// var newProduct response.ProductResponse

	// for _, v := range cart.Products {
	// 	fmt.Println("controller")
	// 	fmt.Println(v.Product.FileUrl)
	// 	newProduct = response.FromUsecase(v.Product)
	// 	newProducts.Product = newProduct
	// 	newProducts.Quantity = v.Quantity
	// 	newProductsList = append(newProductsList, newProducts)
	// }

	return CartResponse{
		Id:         cart.Id,
		UserId:     cart.UserId,
		Products:   cart.Products,
		TotalPrice: cart.TotalPrice,
		CreatedAt:  cart.CreatedAt,
		UpdatedAt:  cart.UpdatedAt,
	}
}

func FromUsecaseList(cart []carts.Cart) []CartResponse {
	var cartResponse []CartResponse

	for _, v := range cart {
		cartResponse = append(cartResponse, FromUsecase(v))
	}

	return cartResponse
}
