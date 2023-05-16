package user

import (
	"errors"

	"github.com/gDenisLit/item-server-go/cmd/services/logger"
	"github.com/gDenisLit/item-server-go/cmd/services/response"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(ctx *fiber.Ctx) error {
	logger.Warn("IP:", ctx.IP(), "requests users")
	users, err := UserService.Query()
	if err != nil {
		return response.ServerError(ctx)
	}
	logger.Info("Sending list of users to IP:", ctx.IP())
	return response.Success(ctx, users)
}

func GetUserById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return response.BadRequest(ctx, errors.New("invalid id"))
	}
	logger.Warn("IP:", ctx.IP(), "request a user with id:", id)
	item, err := UserService.GetById(id)
	if err != nil {
		return response.ServiceError(ctx, "GetUserById", err)
	}
	logger.Info("Sending user to IP:", ctx.IP())
	return response.Success(ctx, item)
}

func RemoveUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return response.BadRequest(ctx, errors.New("invalid id"))
	}
	logger.Warn("IP:", ctx.IP(), "requests to remove user with id:", id)
	itemId, err := UserService.Remove(id)
	if err != nil {
		return response.ServiceError(ctx, "RemoveUser", err)
	}
	logger.Info("User with id:", id, "was removed by:", ctx.IP())
	return response.Success(ctx, itemId)
}
