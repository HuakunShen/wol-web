package middleware

import (
	"fmt"
	"strconv"
	"time"

	"github.com/HuakunShen/golang-auth/controllers"
	"github.com/gofiber/fiber/v2"
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
		userId, err := strconv.ParseUint(fmt.Sprintf("%v", claims["id"]), 10, 32)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "server error",
			})
		}
		ctx.Locals("claims", claims)
		ctx.Locals("username", claims["username"])
		ctx.Locals("id", uint(userId))
		ctx.Locals("uid", claims["id"])
		ctx.Locals("exp", claims["exp"])
		return ctx.Next()
	}
}

