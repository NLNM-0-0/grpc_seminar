package service

import (
	"context"
	"github.com/segmentio/ksuid"
	"product/model"
	"product/repo"
)

type ProductService interface {
	PostProduct(ctx context.Context, name string, price float64) (*model.Product, error)
	GetProduct(ctx context.Context, id string) (*model.Product, error)
	AddRating(ctx context.Context, productId string, score uint64) (*model.Product, error)
}

type productService struct {
	repository repo.ProductRepo
}

func NewService(repository repo.ProductRepo) *productService {
	return &productService{repository}
}

func (service *productService) PostProduct(ctx context.Context, name string, price float64) (*model.Product, error) {
	product := &model.Product{
		Name:  name,
		Price: price,
		Id:    ksuid.New().String(),
	}
	if err := service.repository.CreateProduct(ctx, product); err != nil {
		return nil, err
	}
	return product, nil
}

func (service *productService) GetProduct(ctx context.Context, id string) (*model.Product, error) {
	return service.repository.GetProductById(ctx, id)
}

func (service *productService) AddRating(ctx context.Context, productId string, score uint64) (*model.Product, error) {
	product, err := service.repository.GetProductById(ctx, productId)
	if err != nil {
		return nil, err
	}

	newSum := product.RatingSum + score
	newCount := product.RatingCount + 1

	product.RatingSum = newSum
	product.RatingCount = newCount

	if err := service.repository.SaveProduct(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}
