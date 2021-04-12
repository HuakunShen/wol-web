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
	app.Get("/api/users/count", controllers.UserCount)
	app.Get("/api/computers", middleware.IsAuthenticated(), controllers.GetComputers)
	app.Get("/api/computers/:id", middleware.IsAuthenticated(), controllers.GetComputer)
	app.Post("/api/computers", middleware.IsAuthenticated(), controllers.AddComputer)
	app.Delete("/api/computers/:id", middleware.IsAuthenticated(), controllers.DeleteComputer)
	app.Post("/api/wol/:id", middleware.IsAuthenticated(), controllers.Wol)
}
