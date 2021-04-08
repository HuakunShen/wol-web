package routes

import (
	"github.com/HuakunShen/golang-auth/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/users/register", controllers.Register)
	app.Post("/api/users/login", controllers.Login)
	app.Get("/api/users/user", controllers.User)
	app.Post("/api/users/logout", controllers.Logout)
}
