package products

import (
	"errors"
)

type ProductUseCase struct {
	Repo ProductRepoInterface
}

func NewUseCase(productRepo ProductRepoInterface) ProductUseCaseInterface {
	return &ProductUseCase{
		Repo: productRepo,
	}
}

func (productUseCase *ProductUseCase) AddProduct(product Product) (Product, error) {
	if product.Name == "" {
		return Product{}, errors.New("name cannot be empty")
	}

	if product.Description == "" {
		return Product{}, errors.New("description cannot be empty")
	}

	if product.Price == 0 {
		return Product{}, errors.New("price cannot be empty")
	}

	if product.Stock == 0 {
		return Product{}, errors.New("stock cannot be empty")
	}

	if product.CategoryId == 0 {
		return Product{}, errors.New("category id cannot be empty")
	}

	if product.FileName == "" {
		return Product{}, errors.New("file name cannot be empty")
	}

	productRepo, err := productUseCase.Repo.AddProduct(product)

	if err != nil {
		return Product{}, err
	}

	return productRepo, nil
}

func (productUseCase *ProductUseCase) EditProduct(product Product, id int) (Product, error) {
	if id == 0 {
		return Product{}, errors.New("product ID cannot be empty")
	}

	if product.Name == "" {
		return Product{}, errors.New("name cannot be empty")
	}

	if product.Description == "" {
		return Product{}, errors.New("description cannot be empty")
	}

	if product.Price == 0 {
		return Product{}, errors.New("price cannot be empty")
	}

	if product.Stock == 0 {
		return Product{}, errors.New("stock cannot be empty")
	}

	if product.FileName == "" {
		return Product{}, errors.New("file name cannot be empty")
	}

	if product.CategoryId == 0 {
		return Product{}, errors.New("category id cannot be empty")
	}

	productRepo, err := productUseCase.Repo.EditProduct(product, id)

	if err != nil {
		return Product{}, err
	}

	return productRepo, nil
}

func (productUseCase *ProductUseCase) DeleteProduct(id int) (Product, error) {
	if id == 0 {
		return Product{}, errors.New("product ID cannot be empty")
	}

	productRepo, err := productUseCase.Repo.DeleteProduct(id)

	if err != nil {
		return Product{}, err
	}

	return productRepo, nil
}

func (productUseCase *ProductUseCase) GetProductDetail(id int) (Product, error) {
	if id == 0 {
		return Product{}, errors.New("product ID cannot be empty")
	}

	productRepo, err := productUseCase.Repo.GetProductDetail(id)

	if err != nil {
		return Product{}, err
	}

	return productRepo, nil
}
