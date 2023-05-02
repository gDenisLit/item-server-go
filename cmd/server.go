package cmd

import (
	"os"

	"github.com/gDenisLit/item-server-go/cmd/api/auth"
	"github.com/gDenisLit/item-server-go/cmd/api/item"
	"github.com/gDenisLit/item-server-go/cmd/api/user"
	"github.com/gofiber/fiber/v2"
)

func InitServer() {
	app := fiber.New()

	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendString("OK")
	})

	item.SetItemRoutes(app)
	user.SetUserRoutes(app)
	auth.SetAuthRoutes(app)

	port := os.Getenv("PORT")
	app.Listen(port)
}
