package carts

import (
	"context"
	"errors"
	"florist-gin/business/carts"
	cartsDB "florist-gin/drivers/databases/carts"
	"florist-gin/drivers/databases/cartsproducts"
	"net/url"
	"time"

	minio "github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type CartRepository struct {
	Db          *gorm.DB
	MinioClient *minio.Client
}

func NewCartRepository(database *gorm.DB, minioClient *minio.Client) carts.CartRepoInterface {
	return &CartRepository{
		Db:          database,
		MinioClient: minioClient,
	}
}

func (repo *CartRepository) GetCart(userId int) (carts.Cart, error) {
	var cartDb cartsDB.Cart
	var cartsProductsDB []cartsproducts.CartsProducts
	var productList []carts.Product
	var product carts.Product
	totalPrice := 0

	resultFind := repo.Db.Model(&cartsDB.Cart{}).Where("user_id = ?", userId).First(&cartDb)
	if resultFind.Error != nil {
		return carts.Cart{}, errors.New("cart not found")
	}

	resultPreload := repo.Db.Model(&cartsproducts.CartsProducts{}).Where("cart_id = ?", cartDb.Id).Preload("Product").Find(&cartsProductsDB)

	if resultPreload.Error != nil {
		return carts.Cart{}, errors.New("failed to load cart products")
	}

	ctx := context.Background()
	reqParams := make(url.Values)

	cartUseCase := cartDb.ToUseCase()

	for _, cartsProducts := range cartsProductsDB {
		// Fetch the presigned URL for each product
		presignedURL, err := repo.MinioClient.PresignedGetObject(ctx, "florist", "products/"+cartsProducts.Product.FileName, time.Second*24*60*60, reqParams)

		if err != nil {
			return carts.Cart{}, err
		}

		product.Quantity = cartsProducts.Quantity
		product.Id = cartsProducts.Product.Id
		product.Description = cartsProducts.Product.Description
		product.Name = cartsProducts.Product.Name
		product.Price = cartsProducts.Product.Price
		product.Stock = cartsProducts.Product.Stock
		product.FileUrl = presignedURL.String()
		productList = append(productList, product)
		totalPrice += cartsProducts.Product.Price * cartsProducts.Quantity
	}

	cartUseCase.Products = productList
	cartUseCase.TotalPrice = totalPrice

	return cartUseCase, nil
}
