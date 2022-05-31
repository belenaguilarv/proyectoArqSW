package model

type User struct {
	UserId   int    `gorm:"primaryKey; not null; unsigned; auto_increment"`
	UserName string `gorm:"type:varchar(50); not null; unique"`
	Password string `gorm:"type:varchar(255); not null"`
}

type Users []User
