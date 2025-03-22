package model

type OrderProduct struct {
	Id        string  `json:"id" gorm:"column:id;primaryKey"`
	OrderId   string  `json:"orderId" gorm:"column:order_id"`
	ProductId string  `json:"productId" gorm:"column:product_id"`
	Name      string  `json:"name" gorm:"column:name"`
	Price     float64 `json:"price" gorm:"column:price"`
	Quantity  uint32  `json:"quantity" gorm:"column:quantity"`
}

func (*OrderProduct) TableName() string {
	return "order_products"
}
