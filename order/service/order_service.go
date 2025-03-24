package service

import (
	"context"
	"github.com/segmentio/ksuid"
	"order/model"
	"order/repo"
	"time"
)

type OrderService interface {
	CreateOrder(ctx context.Context, userId string, products []model.OrderProduct) (*model.Order, error)
}

type orderService struct {
	repository repo.OrderRepo
}

func NewOrderService(repo repo.OrderRepo) *orderService {
	return &orderService{repo}
}

func (service orderService) CreateOrder(
	ctx context.Context,
	userId string,
	products []model.OrderProduct,
) (*model.Order, error) {
	order := &model.Order{
		Id:        ksuid.New().String(),
		CreatedAt: time.Now().UTC(),
		UserId:    userId,
		Products:  products,
	}

	order.TotalPrice = 0.0
	for i := range order.Products {
		order.Products[i].Id = ksuid.New().String()
		order.TotalPrice += order.Products[i].Price * float64(order.Products[i].Quantity)
	}

	err := service.repository.CreateOrder(ctx, *order)
	if err != nil {
		return nil, err
	}
	return order, nil
}
