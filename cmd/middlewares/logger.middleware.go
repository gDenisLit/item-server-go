package middlewares

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/gDenisLit/item-server-go/cmd/services"
)

func Log(ctx *fiber.Ctx) error {
	// Log something for each server reqest
	// services.Logger.Info("Request URL: ", )
	return ctx.Next()
}
