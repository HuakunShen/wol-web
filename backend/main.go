package main

import (
	"fmt"
	"log"
	"os"

	"github.com/HuakunShen/wol-web/backend/database"
	"github.com/HuakunShen/wol-web/backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
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

	app.Static("/", "../frontend/dist")
	routes.Setup(app)

	err = app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}
