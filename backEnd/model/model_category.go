package model

type Category struct {
	Id   int    `gorm:"primaryKey;smallint; unsigned; not null; auto_increment"`
	Name string `gorm:"type:varchar(50); not null"`
}

type categories []Category
