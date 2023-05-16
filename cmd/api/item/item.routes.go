package item

import (
	"github.com/gDenisLit/item-server-go/cmd/api/auth"
	"github.com/gDenisLit/item-server-go/cmd/middlewares"
	"github.com/gofiber/fiber/v2"
)

type middleware func(*fiber.Ctx) error

var log middleware = middlewares.Log
var requireAuth middleware = middlewares.RequireAuth(auth.AuthService)

func RegisterRoutes(app *fiber.App) {
	router := app.Group("/api/item")

	router.Get("/", GetItems)
	router.Get("/:id", GetItemById)
	router.Post("/", log, AddItem)
	router.Put("/", log, requireAuth, UpdateItem)
	router.Delete("/:id", log, requireAuth, RemoveItem)
}
