package controllers

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/HuakunShen/wol-web/backend/database"
	"github.com/HuakunShen/wol-web/backend/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *fiber.Ctx) error {
	if os.Getenv("ENABLE_SIGNUP") == "true" {
		// keep going the rest of the signup logic
	} else if os.Getenv("ENABLE_SIGNUP") == "false" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Signup function disabled for this app",
			"error": "Signup function disabled for this app",
		})
	} else {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "invalid environment variable ENABLE_SIGNUP, signup disabled",
			"error": "invalid environment variable ENABLE_SIGNUP, signup disabled",
		})
	}

	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	if len(data["username"]) < 1 || len(data["password"]) < 1 {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"message": "invalid username or password",
			"error": "invalid username or password",
		})
	}
	user := models.User{
		Username: data["username"],
		Password: password,
	}
	if err := database.DB.Create(&user).Error; err != nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Username Exists, Try Another One",
			"error": err,
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data": user,
	})
}

func CreateToken(id uint, username string, validMinutes uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(validMinutes))
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return t, err
	}
	return t, nil
}

func ParseCookieJWT(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Fail to parse JWT")
	}
	return claims, nil
}

func Login(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "fail to parse",
			"error": err,
		})
	}

	var user models.User
	if err := database.DB.Where("username = ?", data["username"]).First(&user).Error; err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
			"error": err,
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		ctx.Status(fiber.StatusForbidden)
		return ctx.JSON(fiber.Map{
			"message": "incorrect password",
			"error": err,
		})
	}
	validTimeMinutes := os.Getenv("JWT_VALID_TIME")
	validMinutes, err := strconv.ParseInt(validTimeMinutes, 10, 32)
	if err != nil {
		panic(err)
	}

	token, err := CreateToken(user.Id, user.Username, uint(validMinutes))

	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.JSON(fiber.Map{
			"message": "could not login, server error",
			"error": err,
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Minute * time.Duration(validMinutes)),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data": user,
	})
}

func User(ctx *fiber.Ctx) error {
	tokenString := ctx.Cookies("jwt")
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		ctx.Status(fiber.StatusUnauthorized)
		fmt.Println(err)
		return ctx.JSON(fiber.Map{
			"message": "unauthenticated",
			"error": err,
		})
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Server Error",
		})
	}
	var user models.User
	if err := database.DB.Where("username = ?", claims["username"]).First(&user).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User Not Authenticated",
			"error": err,
		})
	}
	return ctx.JSON(fiber.Map{"data": user, "message": "success",})
}

func Logout(ctx *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}
