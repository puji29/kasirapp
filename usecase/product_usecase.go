package usecase

import (
	"fmt"
	"main/model"
	"main/repository"
)

type ProductUsecase interface {
	RegisterNewProduct(payload model.Product) (model.Product, error)
	FindAllProduct() ([]model.Product, error)
}

type productUsecase struct {
	repo repository.ProductRepository
}

// FindAllProduct implements ProductUsecase.
func (u *productUsecase) FindAllProduct() ([]model.Product, error) {
	return u.repo.LIst()
}

func (u *productUsecase) RegisterNewProduct(payload model.Product) (model.Product, error) {
	product, err := u.repo.Create(payload)
	if err != nil {
		return model.Product{}, fmt.Errorf("oops, failed to save data product:", err.Error())
	}
	return product, nil
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return &productUsecase{repo: repo}
}
