package controllers

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/HuakunShen/wol-web/backend/database"
	"github.com/HuakunShen/wol-web/backend/middleware"
	"github.com/HuakunShen/wol-web/backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}
	// check if one user exists, if exists disallow sign up
	var count int64
	if err := database.DB.Model(&models.User{}).Count(&count).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error, DB Error",
			"error":   err,
		})
	}

	num_user_allowed, err := strconv.Atoi(os.Getenv("NUM_USER_ALLOWED"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Fail to parse NUM_USER_ALLOWED env variable",
			"error":   err,
		})
	}

	if count >= int64(num_user_allowed) {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": fmt.Sprintf("Not allowed to add more than %d users", num_user_allowed),
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	if len(data["username"]) < 1 || len(data["password"]) < 1 {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"message": "invalid username or password",
			"error":   "invalid username or password",
		})
	}
	user := models.User{
		Username: data["username"],
		Password: password,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Username Exists, Try Another One",
			"error":   err,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserCount(ctx *fiber.Ctx) error {
	num_user_allowed, err := strconv.Atoi(os.Getenv("NUM_USER_ALLOWED"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Fail to parse NUM_USER_ALLOWED env variable",
			"error":   err,
		})
	}
	var count int64
	if err := database.DB.Model(&models.User{}).Count(&count).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error, DB Error",
			"error":   err,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data": fiber.Map{
			"num_user_allowed": num_user_allowed,
			"user_count":       count,
		},
	})
}

func CreateToken(id uint, username string, validMinutes uint) (string, error) {
	expirationTime := time.Now().Add(time.Minute * time.Duration(validMinutes))
	claims := &middleware.AuthClaims{
		Username: username,
		Id:       strconv.Itoa(int(id)),
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}

func Login(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Fail to Parse Request Body",
			"error":   err,
		})
	}

	var user models.User
	if err := database.DB.Where("username = ?", data["username"]).First(&user).Error; err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User Not Found",
			"error":   err,
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		ctx.Status(fiber.StatusForbidden)
		return ctx.JSON(fiber.Map{
			"message": "Incorrect Password",
			"error":   err,
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
			"message": "Could not Login, Server Error related to JWT",
			"error":   err,
		})
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Minute * time.Duration(validMinutes)),
		HTTPOnly: false,
	}

	ctx.Cookie(&cookie)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    user,
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
			"error":   err,
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
			"error":   err,
		})
	}
	return ctx.JSON(fiber.Map{"data": user, "message": "success"})
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
