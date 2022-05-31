package model

type Product struct {
	ProductId        int     `gorm:"primaryKey; smallint; not null; unsigned; auto_increment"`
	ProductName      string  `gorm:"type:varchar(350); not null; unique"`
	ProductPicture   string  `gorm:"type:varchar(350); not null"`
	ProductUnitPrice float32 `gorm:"type:smallint; unsigned; not null"`
}

type ProductsDto []Product
