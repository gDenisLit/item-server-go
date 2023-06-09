package auth

import (
	"github.com/gDenisLit/item-server-go/cmd/middlewares"
	"github.com/gofiber/fiber/v2"
)

type middleware func(*fiber.Ctx) error

var log middleware = middlewares.Log

func SetAuthRoutes(app *fiber.App) {
	router := app.Group("/api/auth")

	router.Post("/login", log, Login)
	router.Post("/signup", log, Signup)
	router.Post("/logout", log, Logout)
}
