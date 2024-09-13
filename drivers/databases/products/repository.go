package products

import (
	"bytes"
	"context"
	"errors"
	"florist-gin/business/products"
	"net/url"
	"time"

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

	file := bytes.NewReader(product.File)

	_, err := repo.MinioClient.PutObject(ctx, "florist", "products/"+product.FileName, file, -1, minio.PutObjectOptions{
		ContentType: "image/jpeg",
	})

	if err != nil {
		return products.Product{}, err
	}

	result := repo.Db.Create(&productDB)

	if result.Error != nil {
		return products.Product{}, result.Error
	}

	return productDB.ToUsecase(), nil
}

func (repo *ProductRepository) EditProduct(product products.Product, id int) (products.Product, error) {
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

	errRemove := repo.MinioClient.RemoveObject(ctx, "florist", "products/"+newProduct.FileName, minio.RemoveObjectOptions{})

	if errRemove != nil {
		return products.Product{}, errRemove
	}

	file := bytes.NewReader(product.File)

	_, err := repo.MinioClient.PutObject(ctx, "florist", "products/"+productDb.FileName, file, -1, minio.PutObjectOptions{
		ContentType: "image/jpeg",
	})

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

func (repo *ProductRepository) DeleteProduct(id int) (products.Product, error) {
	var productDb Product

	resultFind := repo.Db.First(&productDb, id)

	if resultFind.Error != nil {
		return products.Product{}, errors.New("Product not found")
	}

	ctx := context.Background()

	errRemove := repo.MinioClient.RemoveObject(ctx, "florist", "products/"+productDb.FileName, minio.RemoveObjectOptions{})

	if errRemove != nil {
		return products.Product{}, errRemove
	}

	result := repo.Db.Delete(&productDb, id)

	if result.Error != nil {
		return products.Product{}, errors.New("Product not found")
	}

	return productDb.ToUsecase(), nil
}

func (repo *ProductRepository) GetProductDetail(id int) (products.Product, error) {
	var productDb Product

	resultFind := repo.Db.Preload("Category").First(&productDb, id)

	if resultFind.Error != nil {
		return products.Product{}, errors.New("product not found")
	}

	ctx := context.Background()

	// Set request parameters for content-disposition.
	reqParams := make(url.Values)
	// reqParams.Set("response-content-disposition", "inline")

	// Generates a presigned url which expires in a day.
	presignedURL, err := repo.MinioClient.PresignedGetObject(ctx, "florist", "products/"+productDb.FileName, time.Second*24*60*60, reqParams)

	if err != nil {
		return products.Product{}, err
	}

	productUC := productDb.ToUsecase()

	productUC.FileUrl = presignedURL

	return productUC, nil
}

func (repo *ProductRepository) GetAllProduct(categoryId *int) ([]products.Product, error) {
	var newProducts []Product

	query := repo.Db.Preload("Category")

	// Conditionally add the filter for categoryId if it's provided
	if categoryId != nil {
		query = query.Where("category_id = ?", *categoryId)
	}

	result := query.Find(&newProducts)

	if result.Error != nil {
		return []products.Product{}, errors.New("product not found")
	}

	ctx := context.Background()

	// Set request parameters for content-disposition.
	reqParams := make(url.Values)
	// reqParams.Set("response-content-disposition", "inline")

	// Generates a presigned url which expires in a day.
	productUCs := make([]products.Product, len(newProducts))

	for i, product := range newProducts {
		presignedURL, err := repo.MinioClient.PresignedGetObject(ctx, "florist", "products/"+product.FileName, time.Second*24*60*60, reqParams)

		if err != nil {
			return []products.Product{}, err
		}

		productUCs[i] = product.ToUsecase()

		productUCs[i].FileUrl = presignedURL
	}

	return productUCs, nil
}
