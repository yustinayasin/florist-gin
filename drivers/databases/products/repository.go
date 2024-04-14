package products

import (
	"context"
	"errors"
	"florist-gin/business/products"

	minio "github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type ProductRepository struct {
	Db          *gorm.DB
	MinioClient *minio.Client
}

func NewProductRepository(database *gorm.DB, minioClient *minio.Client) products.ProductRepoInterface {
	return &ProductRepository{
		Db:          database,
		MinioClient: minioClient,
	}
}

func (repo *ProductRepository) AddProduct(product products.Product) (products.Product, error) {
	productDB := FromUsecase(product)

	ctx := context.Background()

	_, err := repo.MinioClient.PutObject(ctx, "florist", "florist/products/"+product.FileName, product.File, -1, minio.PutObjectOptions{})

	if err != nil {
		return products.Product{}, err
	}

	result := repo.Db.Create(&productDB)

	if result.Error != nil {
		return products.Product{}, result.Error
	}

	return productDB.ToUsecase(), nil
}

func (repo *ProductRepository) EditProduct(product products.Product, id uint32) (products.Product, error) {
	productDb := FromUsecase(product)

	var newProduct Product

	result := repo.Db.First(&newProduct, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return products.Product{}, errors.New("Product not found")
		}
		return products.Product{}, errors.New("error in database")
	}

	ctx := context.Background()

	errRemove := repo.MinioClient.RemoveObject(ctx, "florist", "florist/products/"+newProduct.FileName, minio.RemoveObjectOptions{})

	if errRemove != nil {
		return products.Product{}, errRemove
	}

	_, err := repo.MinioClient.PutObject(ctx, "florist", "florist/products/"+productDb.FileName, product.File, -1, minio.PutObjectOptions{})

	if err != nil {
		return products.Product{}, err
	}

	newProduct.Name = productDb.Name
	newProduct.Description = productDb.Description
	newProduct.Price = productDb.Price
	newProduct.Stock = productDb.Stock
	newProduct.FileName = productDb.FileName
	newProduct.CategoryId = productDb.CategoryId

	repo.Db.Save(&newProduct)
	return newProduct.ToUsecase(), nil
}

func (repo *ProductRepository) DeleteProduct(id uint32) (products.Product, error) {
	var productDb Product

	resultFind := repo.Db.First(&productDb, id)

	if resultFind.Error != nil {
		return products.Product{}, errors.New("Product not found")
	}

	ctx := context.Background()

	errRemove := repo.MinioClient.RemoveObject(ctx, "florist", "florist/products/"+productDb.FileName, minio.RemoveObjectOptions{})

	if errRemove != nil {
		return products.Product{}, errRemove
	}

	result := repo.Db.Delete(&productDb, id)

	if result.Error != nil {
		return products.Product{}, errors.New("Product not found")
	}

	return productDb.ToUsecase(), nil
}

func (repo *ProductRepository) GetProductDetail(id uint32) (products.Product, error) {
	var productDb Product

	resultFind := repo.Db.Preload("Category").First(&productDb, id)

	if resultFind.Error != nil {
		return products.Product{}, errors.New("Product not found")
	}

	ctx := context.Background()

	object, err := repo.MinioClient.GetObject(ctx, "florist", "florist/products/"+productDb.FileName, minio.GetObjectOptions{})

	if err != nil {
		return products.Product{}, err
	}

	productUC := productDb.ToUsecase()

	productUC.File = object

	return productUC, nil
}
