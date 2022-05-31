package model

type Order struct {
	OrderId int `gorm:"primaryKey;smallint; unsigned; not null; auto_increment"`
	//OrderDate  date `gorm:"type: date"`
	TotalPrice int `gorm:"type:smallint; unsigned; not null"`
	Quantity   int `gorm:"type:smallint; unsigned; not null"`

	ProductId int `gorm:"primaryKey; smallint; not null; unsigned; auto_increment"`
}

type Orders []Order
