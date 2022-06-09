package model

type Product struct {
	Id          int     `gorm:"primaryKey; smallint; not null; unsigned; auto_increment"`
	Name        string  `gorm:"type:varchar(350); not null; unique"`
	Description string  `gorm:"type:varchar(350); not null"`
	Picture     string  `gorm:"type:varchar(350); not null"`
	Price       float32 `gorm:"type:float; not null"`
	Category    string  `gorm:"type:varchar(350); not null"`
	stock       int     `gorm:"type:smallint; unsigned; not null"`
}

type Products []Product
