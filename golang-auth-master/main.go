package main

import (
	"fmt"
	"github.com/HuakunShen/golang-auth/database"
	"github.com/HuakunShen/golang-auth/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	port := os.Getenv("PORT")
	fmt.Println("PORT:", port)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	err = app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}
