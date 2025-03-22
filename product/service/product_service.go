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
	GetProductByIds(ctx context.Context, ids []string) ([]*model.Product, error)
}

type productService struct {
	repository repo.ProductRepo
}

func NewService(repo repo.ProductRepo) *productService {
	return &productService{repo}
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

func (service *productService) GetProductByIds(ctx context.Context, ids []string) ([]*model.Product, error) {
	return service.repository.GetProductByIds(ctx, ids)
}
