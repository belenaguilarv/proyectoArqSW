package model

type Product struct {
	ProductId   int    `gorm:"primaryKey; smallint; not null; unsigned; auto_increment"`
	Name        string `gorm:"type:varchar(350); not null; unique"`
	Description string `gorm:"type:varchar(350); not null"`
	//Picture             // no se que estructura usar ... ????
	Price float32 `gorm:"type:float; not null"`
}

type Products []Product
