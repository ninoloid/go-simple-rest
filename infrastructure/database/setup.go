package database

import (
	"github.com/ninoloid/go-simple-rest/common/helper"
	users "github.com/ninoloid/go-simple-rest/domain/users/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:user@tcp(localhost:3306)/goSimpleRest"))
	helper.PanicIfError(err)

	database.AutoMigrate(&users.User{})

	DB = database
}
