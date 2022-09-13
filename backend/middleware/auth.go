package middleware

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type AuthClaims struct {
	Username string `json:"username"`
	Id       string `json:"id"`
	jwt.StandardClaims
}

func IsAuthenticated() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		tokenString := ctx.Cookies("jwt")
		claims := &AuthClaims{}
		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		fmt.Println(err)

		if err != nil {
			fmt.Println(err)
			if err == jwt.ErrSignatureInvalid {
				return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Session Expired",
					"error":   err,
				})
			}
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid Token",
				"error":   err,
			})
		}
		if !tkn.Valid {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Session Expired",
			})
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%v", claims.Id), 10, 32)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "server error",
			})
		}
		ctx.Locals("claims", claims)
		ctx.Locals("username", claims.Username)
		ctx.Locals("id", uint(userId))
		ctx.Locals("uid", claims.Id)
		// ctx.Locals("exp", claims["exp"])
		return ctx.Next()
	}
}
