package controllers

import (
	"github.com/HuakunShen/golang-auth/database"
	"github.com/HuakunShen/golang-auth/models"
	"github.com/gofiber/fiber/v2"
)

func AddMac(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	mac := models.Mac{
		Name: data["name"],
		Mac:  data["mac"],
	}
	// TODO: Verify mac name doesn't exist, if exists, set 400 as status
	database.DB.Create(&mac)

	return c.JSON(mac)
}


