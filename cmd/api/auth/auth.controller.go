package auth

import (
	"github.com/gDenisLit/item-server-go/cmd/models"
	"github.com/gDenisLit/item-server-go/cmd/services/response"

	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	credentials := new(models.User)
	parseErr := ctx.BodyParser(credentials)
	validateErr := credentials.ValidateLoginCredentials()

	if parseErr != nil || validateErr != nil {
		return response.ParsingError(ctx, "Login", parseErr, validateErr)
	}
	user, err := AuthService.Login(credentials)
	if err != nil {
		return response.ServiceError(ctx, "Login", err)
	}
	loginToken, err := AuthService.GetLoginToken(user)
	if err != nil {
		return response.ServerError(ctx)
	}
	return response.LoginSuccess(ctx, loginToken, user.Minify())
}

func Signup(ctx *fiber.Ctx) error {
	credentials := new(models.User)
	parseErr := ctx.BodyParser(credentials)
	validateErr := credentials.ValidateLoginCredentials()

	if parseErr != nil || validateErr != nil {
		return response.ParsingError(ctx, "Signup", parseErr, validateErr)
	}
	user, err := AuthService.Signup(credentials)
	if err != nil {
		return response.ServiceError(ctx, "Signup", err)
	}
	loginToken, err := AuthService.GetLoginToken(user)
	if err != nil {
		return response.ServerError(ctx)
	}
	return response.LoginSuccess(ctx, loginToken, user.Minify())
}

func Logout(ctx *fiber.Ctx) error {
	return response.LogoutSuccess(ctx)
}
