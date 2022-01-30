package database

import (
	"github.com/HuakunShen/wol-web/backend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(sqlite.Open("./data/wol.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = connection
	err = connection.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
	err = connection.AutoMigrate(&models.Computer{})
	if err != nil {
		panic(err)
	}
}
