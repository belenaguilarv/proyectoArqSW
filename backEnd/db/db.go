package db

import (
	userClient "github.com/belenaguilarv/proyectoArqSW/backEnd/clients"
	"github.com/belenaguilarv/proyectoArqSW/backEnd/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {

	dsn := "root:@tcp(127.0.0.1:3306)/nodelogin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	userClient.Db = db
}
func StartDbEngine() {
	db.AutoMigrate(&model.User{}) // crea una tabla en plural de "user" o la usa si esta creada
	log.Info("Finishing Migration Database Tables")
}
