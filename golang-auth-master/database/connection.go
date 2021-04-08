package database

import (
	"fmt"
	"github.com/HuakunShen/golang-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect() {
	connection_string := fmt.Sprint(os.Getenv("DB_USERNAME"), ":", os.Getenv("DB_PASSWORD"), "@/", os.Getenv("DB_DATABASE"))
	fmt.Println("connection_string: " + connection_string)
	connection, err := gorm.Open(mysql.Open(connection_string), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = connection

	err = connection.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
}
