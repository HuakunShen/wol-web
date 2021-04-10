package controllers

import (
	"fmt"

	"github.com/HuakunShen/wol-web/backend/database"
	"github.com/HuakunShen/wol-web/backend/models"
	wol "github.com/HuakunShen/wol/wol-go"
	"github.com/gofiber/fiber/v2"
)

func Wol(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	userId := ctx.Locals("id")
	var mac models.Mac
	if err := database.DB.Where("id = ?", id).Where("user_id = ?", userId).First(&mac).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Mac Not Found or you don't have permission",
			"error":   fmt.Sprintf("%v", err),
		})
	}
	// todo add ip and port
	if err := wol.WakeOnLan(mac.Mac, mac.Ip, id); err != nil {
		fmt.Println("Error:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "fail to wol",
			"error":   err,
		})
	}
	return ctx.JSON(fiber.Map{"message": "success"})
}
