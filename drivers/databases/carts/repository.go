package carts

import (
	"errors"
	"florist-gin/business/carts"

	"gorm.io/gorm"
)

type CartRepository struct {
	Db *gorm.DB
}

func NewCartRepository(database *gorm.DB) carts.CartRepoInterface {
	return &CartRepository{
		Db: database,
	}
}

func (repo *CartRepository) GetCart(id int) (carts.Cart, error) {
	var cartDb Cart

	resultFind := repo.Db.Model(&Cart{}).Preload("Product").First(&cartDb, id)

	if resultFind.Error != nil {
		return carts.Cart{}, errors.New("cart not found")
	}

	return cartDb.ToUseCase(), nil
}
