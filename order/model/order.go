package model

import "time"

type Order struct {
	Id         string         `json:"id" gorm:"column:id;primaryKey"`
	CreatedAt  time.Time      `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	UserId     string         `json:"userId" gorm:"column:user_id"`
	TotalPrice float64        `json:"totalPrice" gorm:"column:total_price"`
	Products   []OrderProduct `json:"products" gorm:"foreignKey:OrderId;constraint:OnDelete:CASCADE"`
}

func (*Order) TableName() string {
	return "orders"
}
