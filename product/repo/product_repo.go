package repo

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"product/model"
	"product/utils"
)

type ProductRepo interface {
	CreateProduct(ctx context.Context, data *model.Product) error
	GetProductById(ctx context.Context, productId string) (*model.Product, error)
	GetProductByIds(ctx context.Context, productIds []string) ([]*model.Product, error)
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *productRepo {
	return &productRepo{db: db}
}

func (repo *productRepo) CreateProduct(ctx context.Context, data *model.Product) error {
	db := repo.db

	if err := db.Create(data).Error; err != nil {
		return utils.ErrDB(err)
	}

	return nil
}

func (repo *productRepo) GetProductById(ctx context.Context, productId string) (*model.Product, error) {
	var data model.Product
	db := repo.db

	if err := db.Where(map[string]interface{}{"id": productId}).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.ErrRecordNotFound()
		}
		return nil, utils.ErrDB(err)
	}

	return &data, nil
}

func (repo *productRepo) GetProductByIds(ctx context.Context, productIds []string) ([]*model.Product, error) {
	var products []*model.Product
	db := repo.db

	if err := db.Where("id IN ?", productIds).Find(&products).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.ErrRecordNotFound()
		}
		return nil, utils.ErrDB(err)
	}

	return products, nil
}
