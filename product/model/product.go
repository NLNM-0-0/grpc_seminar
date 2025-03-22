package model

type Product struct {
	Id    string  `json:"id" gorm:"column:id;primaryKey"`
	Name  string  `json:"name" gorm:"column:name"`
	Price float64 `json:"price" gorm:"column:price"`
}

func (*Product) TableName() string {
	return "products"
}
