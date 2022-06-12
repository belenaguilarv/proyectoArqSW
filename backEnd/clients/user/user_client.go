package user

import (
	"github.com/belenaguilarv/proyectoArqSW/backEnd/dto"
	"github.com/belenaguilarv/proyectoArqSW/backEnd/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetUserById(id int) model.User {
	var user model.User
	Db.Where("id = ?", id).First(&user)
	log.Debug("User: ", user)
	return user
}

func GetUsers() model.Users {
	var users model.Users
	Db.Find(&users)
	log.Debug("Users: ", users)
	return users
}

func GetUserByName(LoginDto dto.LoginDto) model.User {
	var user model.User
	Db.First(&user, "Name = ? AND password = ?", LoginDto.Name, LoginDto.Password)
	log.Debug("User: ", user)
	return user
}
