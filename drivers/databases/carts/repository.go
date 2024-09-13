package carts

import (
	"context"
	"errors"
	"florist-gin/business/carts"
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
	var cartDb Cart

	resultFind := repo.Db.Model(&Cart{}).Preload("Product").Where("user_id = ?", userId).First(&cartDb)

	if resultFind.Error != nil {
		return carts.Cart{}, errors.New("cart not found")
	}

	ctx := context.Background()
	reqParams := make(url.Values)

	cartUseCase := cartDb.ToUseCase()

	for i, product := range cartUseCase.Product {
		presignedURL, err := repo.MinioClient.PresignedGetObject(ctx, "florist", "products/"+product.FileName, time.Second*24*60*60, reqParams)

		if err != nil {
			return carts.Cart{}, err
		}

		cartUseCase.Product[i].FileUrl = presignedURL
	}

	return cartUseCase, nil
}
