package auth

import (
	"github.com/gDenisLit/item-server-go/cmd/dtos"
	"github.com/gDenisLit/item-server-go/cmd/models"

	// "github.com/gDenisLit/item-server-go/cmd/services"
	"github.com/gofiber/fiber/v2"
)

func OnLogin(ctx *fiber.Ctx) error {

	credentials := new(dtos.LoginDTO)
	err := ctx.BodyParser(credentials)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": "Invalid credentials object",
		})
	}

	// services.Log.Info("Login request", credentials)
	user, err := Login(credentials)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}
	return sendMiniUser(ctx, user)
}

func OnSignup(ctx *fiber.Ctx) error {

	credentials := new(dtos.SignupDTO)
	err := ctx.BodyParser(credentials)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": "Invalid credentials object",
		})
	}

	// services.Log.Info("Signup request", credentials)
	user, err := Signup(credentials)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return sendMiniUser(ctx, user)
}

func sendMiniUser(ctx *fiber.Ctx, user *models.User) error {
	miniUser := getMiniUser(user)
	loginToken, err := GetLoginToken(miniUser)
	if err != nil {
		// services.Log.Debug("error in auth controller:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	// services.Log.Info("User logged in:", miniUser)
	ctx.Cookie(getCookie(&loginToken))
	return ctx.Status(fiber.StatusOK).JSON(miniUser)
}

func OnLogout(ctx *fiber.Ctx) error {
	return nil
}

func getMiniUser(loggedinUser *models.User) *dtos.UserDTO {
	miniUser := dtos.UserDTO{
		ID:       loggedinUser.ID,
		Username: loggedinUser.Username,
		Fullname: loggedinUser.Fullname,
		ImgUrl:   loggedinUser.ImgUrl,
	}
	return &miniUser
}

func getCookie(loginToken *string) *fiber.Cookie {
	cookie := &fiber.Cookie{
		Name:     "loginToken",
		Value:    *loginToken,
		Secure:   true,
		SameSite: "None",
	}
	return cookie
}
