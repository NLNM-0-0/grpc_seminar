package repo

import (
	"context"
	"gorm.io/gorm"
	"order/model"
	"order/utils"
)

type OrderRepo interface {
	CreateOrder(ctx context.Context, data model.Order) error
	GetOrdersForUser(ctx context.Context, userId string) ([]model.Order, error)
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *orderRepo {
	return &orderRepo{db: db}
}

func (repo *orderRepo) CreateOrder(ctx context.Context, data model.Order) error {
	db := repo.db

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&data).Error; err != nil {
			return err
		}

		if len(data.Products) > 0 {
			for i := range data.Products {
				data.Products[i].OrderId = data.Id
			}
			if err := tx.Create(&data.Products).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return utils.ErrDB(err)
	}

	return nil
}

func (repo *orderRepo) GetOrdersForUser(ctx context.Context, userId string) ([]model.Order, error) {
	var orders []model.Order

	err := repo.db.
		Where("user_id = ?", userId).
		Preload("order_products").
		Find(&orders).Error

	if err != nil {
		return nil, utils.ErrDB(err)
	}

	return orders, nil
}
