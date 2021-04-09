package middleware

import (
	"fmt"
	"github.com/HuakunShen/golang-auth/controllers"
	"github.com/gofiber/fiber/v2"
	"time"
)

func IsAuthenticated() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		tokenString := ctx.Cookies("jwt")
		claims, err := controllers.ParseCookieJWT(tokenString)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "unauthorized",
			})
		}
		expDate, err := time.Parse(time.RFC3339, fmt.Sprintf("%v", claims["exp"]))
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Not Authorized",
			})
		}
		if expDate.Before(time.Now()) {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Session Expired",
			})
		}
		ctx.Locals("claims", claims)
		return ctx.Next()
	}
}

