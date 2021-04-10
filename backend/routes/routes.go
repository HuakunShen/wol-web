package routes

import (
	"github.com/HuakunShen/wol-web/backend/controllers"
	"github.com/HuakunShen/wol-web/backend/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup(app *fiber.App) {
	app.Use(logger.New())
	app.Post("/api/users/register", controllers.Register)
	app.Post("/api/users/login", controllers.Login)
	app.Get("/api/users/user", middleware.IsAuthenticated(), controllers.User)
	app.Post("/api/users/logout", controllers.Logout)
	app.Get("/api/mac", middleware.IsAuthenticated(), controllers.GetMacs)
	app.Get("/api/mac/:id", middleware.IsAuthenticated(), controllers.GetMac)
	app.Post("/api/mac", middleware.IsAuthenticated(), controllers.AddMac)
	app.Delete("/api/mac/:id", middleware.IsAuthenticated(), controllers.DeleteMac)
	app.Post("/api/wol/:id", middleware.IsAuthenticated(), controllers.Wol)
}
