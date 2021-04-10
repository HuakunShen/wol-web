package database

import (
	"fmt"
	"os"

	"github.com/HuakunShen/wol-web/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprint("host=" + os.Getenv("POSTGRES_HOST"), " user=" + os.Getenv("POSTGRES_USER"), " password=" + os.Getenv("POSTGRES_PASSWORD"), " dbname=", os.Getenv("POSTGRES_DB"), " port=" + os.Getenv("POSTGRES_PORT"), " sslmode=disable ", "TimeZone=", os.Getenv("POSTGRES_TIMESZONE"),)
	fmt.Println(dsn)
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = connection

	err = connection.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
}
