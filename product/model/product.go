package model

type Product struct {
	Id          string  `json:"id" gorm:"column:id;primaryKey"`
	Name        string  `json:"name" gorm:"column:name;unique"`
	Price       float64 `json:"price" gorm:"column:price"`
	RatingCount uint32  `json:"ratingCount" gorm:"column:rating_count"`
	RatingSum   uint64  `json:"ratingSum" gorm:"column:rating_sum"`
}

func (*Product) TableName() string {
	return "products"
}
