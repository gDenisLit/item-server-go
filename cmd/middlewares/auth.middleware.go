package middlewares

import (
	"github.com/gDenisLit/item-server-go/cmd/dtos"
	"github.com/gDenisLit/item-server-go/cmd/services"
	"github.com/gofiber/fiber/v2"
)

func RequireAuth(ctx *fiber.Ctx) error {
	loginToken := ctx.Cookies("loginToken")
	if loginToken == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	userJson := &dtos.LoginDTO{}
	err := services.Decode("loginToken", loginToken, userJson)
	if err != nil {
		// services.Log.Debug("error decoding token:", err)
		return err
	}
	return ctx.Next()
}
