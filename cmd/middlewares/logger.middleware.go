package middlewares

import (
	"github.com/gDenisLit/item-server-go/cmd/services/logger"
	"github.com/gofiber/fiber/v2"
)

func Log(ctx *fiber.Ctx) error {
	logger.Info(ctx.Method(), "Request from", ctx.IP(), "for", ctx.Path())
	return ctx.Next()
}
