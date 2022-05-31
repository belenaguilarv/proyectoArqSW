package model

type User struct {
	Id       int    `gorm:"primaryKey; not null; auto_increment"`
	Name     string `gorm:"type:varchar(50); not null"`
	Password string `gorm:"type:varchar(255); not null"`
}

type Users []User
