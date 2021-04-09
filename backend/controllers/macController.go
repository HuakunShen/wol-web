package controllers

import (
	"fmt"
	"strconv"

	"github.com/HuakunShen/golang-auth/database"
	"github.com/HuakunShen/golang-auth/models"
	"github.com/gofiber/fiber/v2"
)

func AddMac(ctx *fiber.Ctx) error {
	var data map[string]string
	uid := ctx.Locals("id")
	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Fail to parse body",
		})
	}
	userId, err := strconv.ParseUint(fmt.Sprintf("%v", uid), 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Fail to parse userId",
		})
	}
	mac := models.Mac{
		UserId: uint(userId),
		Name: data["name"],
		Mac:  data["mac"],
	}
	if err := database.DB.Create(&mac).Error; err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Mac Exists, Try Another One",
		})
	} else {
		return ctx.Status(fiber.StatusCreated).JSON(mac)
	}
}

func GetMac(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	userId := ctx.Locals("id")
	var mac models.Mac
	if err := database.DB.Where("id = ?", id).Where("user_id = ?", userId).First(&mac).Error; err != nil {
		fmt.Println(err)
	}
	if mac.UserId != userId {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "You have no permission to get this resource",
		})
	}
	return ctx.JSON(fiber.Map{"mac": mac})
}


func GetMacs(ctx *fiber.Ctx) error {
	userId := ctx.Locals("id")
	var macs []models.Mac
	if err := database.DB.Where("user_id = ?", userId).Find(&macs).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Fail to Get Data",
		})
	}
	return ctx.Status(fiber.StatusFound).JSON(fiber.Map{"macs": macs})
}

func DeleteMac(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	userId := ctx.Locals("id")
	var mac models.Mac

	if err := database.DB.Where("id = ?", id).Where("user_id = ?", userId).Delete(&mac); err != nil {
		fmt.Println(err)
		fmt.Println(err==nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Fail to delete, record doesn't Exists or you don't have permission",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"mac": mac,
	})
}
