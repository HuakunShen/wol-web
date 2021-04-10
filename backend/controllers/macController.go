package controllers

import (
	"fmt"
	"strconv"

	"github.com/HuakunShen/wol-web/backend/database"
	"github.com/HuakunShen/wol-web/backend/models"
	"github.com/gofiber/fiber/v2"
)

func AddMac(ctx *fiber.Ctx) error {
	var data map[string]string
	uid := ctx.Locals("id")
	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Fail to parse body",
			"error": err,
		})
	}
	fmt.Println(len(data["mac"]))
	if len(data["mac"]) != 17 && len(data["mac"]) != 12 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Wrong Mac Address Format",
			"error": "Wrong Mac Address Format",
		})
	}
	userId, err := strconv.ParseUint(fmt.Sprintf("%v", uid), 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Fail to parse userId",
			"error": err,
		})
	}
	port, err := strconv.Atoi(data["port"])
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Fail to convert port",
			"error": err,
		})
	}
	mac := models.Mac{
		UserId: uint(userId),
		Name: data["name"],
		Mac:  data["mac"],
		Ip: data["ip"],
		Port:  uint32(port),
	}
	if err := database.DB.Create(&mac).Error; err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Mac Exists, Try Another One",
			"error": err,
		})
	} else {
		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
			"data": mac,
		})
	}
}

func GetMac(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	userId := ctx.Locals("id")
	var mac models.Mac
	if err := database.DB.Where("id = ?", id).Where("user_id = ?", userId).First(&mac).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Mac Not Found or you don't have permission",
			"error": fmt.Sprintf("%v", err),
		})
	}
	return ctx.JSON(fiber.Map{"data": mac})
}

func GetMacs(ctx *fiber.Ctx) error {
	userId := ctx.Locals("id")
	var macs []models.Mac
	if err := database.DB.Where("user_id = ?", userId).Find(&macs).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Fail to Get Data",
			"error": err,
		})
	}
	return ctx.Status(fiber.StatusFound).JSON(fiber.Map{"data": macs})
}

func DeleteMac(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
			"error": err,
		})
	}
	userId := ctx.Locals("id")
	var mac models.Mac

	// verify exist
	if err := database.DB.Where("id = ?", id).Where("user_id = ?", userId).First(&mac).Error; err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data to delete may not exist",
			"error": fmt.Sprintf("%v", err),
		})
	}
	// delete
	if err := database.DB.Where("id = ?", id).Where("user_id = ?", userId).Delete(&mac).Error; err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Fail to delete, record doesn't exists or you don't have permission",
			"error": err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data": mac,
	})
}
