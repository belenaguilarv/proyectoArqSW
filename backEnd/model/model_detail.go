package model

type OrderDetail struct {
	DetailId   int     `gorm:"primaryKey; int; not null; auto_increment"`
	Quantity   int     `gorm:"type:int; not null"`
	Price      float32 `gorm:"type:float; not null"`
	TotalPrice float32 `gorm:"type:float; not null"`
	ProductId  int     `gorm:"type:int; not null"`
	OrderId    int     `gorm:"type:smallint; unsigned; not null"`
}

type OrderDetails []OrderDetail
