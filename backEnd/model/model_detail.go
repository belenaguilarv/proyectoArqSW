package model

type OrderDetail struct {
	DetailId  int     `gorm:"primaryKey; smallint; not null; unsigned; auto_increment"`
	Quantity  int     `gorm:"type:smallint; unsigned; not null"`
	Price     float32 `gorm:"type:float; not null"`
	ProductId int     `gorm:"type:smallint; unsigned; not null"`
	OrderId   int     `gorm:"type:smallint; unsigned; not null"`
}

type OrderDetails []OrderDetail
