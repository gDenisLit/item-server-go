package middlewares

import (
	"github.com/gDenisLit/item-server-go/cmd/models"
	"github.com/gDenisLit/item-server-go/cmd/services/logger"
	"github.com/gDenisLit/item-server-go/cmd/services/response"
	"github.com/gofiber/fiber/v2"
)

type Auth interface {
	ValidateToken(tokne string) (*models.User, error)
}

func RequireAuth(AuthService Auth) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		loginToken := ctx.Cookies("loginToken")
		if loginToken == "" {
			return response.Unauthorized(ctx)
		}

		_, err := AuthService.ValidateToken(loginToken)
		if err != nil {
			logger.Debug("[AuthMiddleware]", err)
			logger.Warn("[AuthMiddleware] Invalid login token:", loginToken, "from IP:", ctx.IP())
			return response.Unauthorized(ctx)
		}
		return ctx.Next()
	}
}
