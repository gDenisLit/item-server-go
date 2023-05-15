package cmd

import (
	"github.com/gDenisLit/item-server-go/cmd/api/auth"
	"github.com/gDenisLit/item-server-go/cmd/api/item"
	"github.com/gDenisLit/item-server-go/cmd/api/user"
	"github.com/gDenisLit/item-server-go/cmd/config"
	"github.com/gofiber/fiber/v2"
)

func InitServer() {
	app := fiber.New()

	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendString("OK")
	})

	item.RegisterRoutes(app)
	user.SetUserRoutes(app)
	auth.SetAuthRoutes(app)

	port := ":" + config.PORT
	app.Listen(port)
}
