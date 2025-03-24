package repo

import (
	"context"
	"gorm.io/gorm"
	"order/model"
	"order/utils"
)

type OrderRepo interface {
	CreateOrder(ctx context.Context, data model.Order) error
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *orderRepo {
	return &orderRepo{db: db}
}

func (repo *orderRepo) CreateOrder(ctx context.Context, data model.Order) error {
	db := repo.db

	if err := db.Create(&data).Error; err != nil {
		return utils.ErrDB(err)
	}

	return nil
}
