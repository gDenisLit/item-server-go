package user

import (
	"github.com/gDenisLit/item-server-go/cmd/services"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(ctx *fiber.Ctx) error {

	services.Log.Info("Getting users")
	users, err := Query()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(users)
}

func GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	services.Log.Info("Getting user with id:", id)
	item, err := GetById(id)

	if err != nil {
		services.Log.Error("User controller Error:", err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(item)
}
