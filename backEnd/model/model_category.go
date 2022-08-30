package model

type Category struct {
	Id   int    `gorm:"primaryKey; int; not null; unsigned; auto_increment"`
	name string `gorm:"type:varchar(255); not null"`
}
