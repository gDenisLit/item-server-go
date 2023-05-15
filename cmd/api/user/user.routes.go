package user

import (
	"github.com/gDenisLit/item-server-go/cmd/middlewares"
	"github.com/gofiber/fiber/v2"
)

type middleware func(*fiber.Ctx) error

var log middleware = middlewares.Log

func SetUserRoutes(app *fiber.App) {
	router := app.Group("/api/user")

	router.Get("/", log, GetUsers)
	router.Get("/:id", log, GetUserById)
	router.Delete("/:id", log, RemoveUser)
}
