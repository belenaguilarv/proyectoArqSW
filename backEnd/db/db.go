package db

import (
	orderClient "github.com/belenaguilarv/proyectoArqSW/backEnd/clients/order"
	orderDetailClient "github.com/belenaguilarv/proyectoArqSW/backEnd/clients/orderDetail"
	productClient "github.com/belenaguilarv/proyectoArqSW/backEnd/clients/product"
	userClient "github.com/belenaguilarv/proyectoArqSW/backEnd/clients/user"
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

	dsn := "root:secret@tcp(127.0.0.1:33060)/nodelogin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	userClient.Db = db
	orderClient.Db = db
	productClient.Db = db
	orderDetailClient.Db = db
}
func StartDbEngine() {

	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.User{}) // crea una tabla en plural de "user" o la usa si esta creada
	db.AutoMigrate(&model.Order{})
	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.OrderDetail{})

	log.Info("Finishing Migration Database Tables")
}
