package middlewares

import "github.com/gofiber/fiber/v2"

func Log(ctx *fiber.Ctx) error {
	// Log something for each server reqest
	return ctx.Next()
}
