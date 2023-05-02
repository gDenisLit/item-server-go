package item

import (
	"github.com/gDenisLit/item-server-go/cmd/middlewares"
	"github.com/gofiber/fiber/v2"
)

type middleware func(*fiber.Ctx) error

var log middleware = middlewares.Log
var auth middleware = middlewares.RequireAuth

func SetItemRoutes(app *fiber.App) {
	router := app.Group("/api/item")

	router.Get("/", log, GetItems)
	router.Get("/:id", log, GetItemById)
	router.Post("/", log, AddItem)
	router.Put("/", log, UpdateItem)
	router.Delete("/:id", log, RemoveItem)
}
