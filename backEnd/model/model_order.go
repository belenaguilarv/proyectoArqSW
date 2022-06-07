package model

type Order struct {
	Id         int     `gorm:"primaryKey;smallint; unsigned; not null; auto_increment"`
	Date       string  `gorm:"type:varchar(50); not null"`
	TotalPrice float32 `gorm:"type:float; not null"`
	Quantity   int     `gorm:"type:smallint; unsigned; not null"`
	UserId     int     `gorm:"type:smallint; unsigned; not null"`
}

type Orders []Order

/*
type OrderWithDetails struct {
	OrderId    int     `gorm:"primaryKey;smallint; unsigned; not null; auto_increment"`
	Date       string  `gorm:"type:varchar(50); not null"`
	TotalPrice float32 `gorm:"type:float; not null"`
	Quantity   int     `gorm:"type:smallint; unsigned; not null"`
	//Details    Details `gorm:"foreignKey:OrderId"` // ???
}
*/
