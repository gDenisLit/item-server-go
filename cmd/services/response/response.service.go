package response

import (
	"errors"

	"github.com/gDenisLit/item-server-go/cmd/models"
	"github.com/gDenisLit/item-server-go/cmd/services/logger"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Success(ctx *fiber.Ctx, data any) error {
	response := makeResponse(200, data, "success")
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func Created(ctx *fiber.Ctx, data any) error {
	response := makeResponse(201, data, "success")
	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func LoginSuccess(ctx *fiber.Ctx, token string, data any) error {
	// TODO
	response := makeResponse(201, data, "success")
	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func LogoutSuccess(ctx *fiber.Ctx) error {
	return nil
}

func BadRequest(ctx *fiber.Ctx, err error) error {
	response := makeResponse(400, nil, "Bad request: "+err.Error())
	return ctx.Status(fiber.StatusBadRequest).JSON(response)
}

func Unauthorized(ctx *fiber.Ctx) error {
	response := makeResponse(401, nil, "Unauthorized")
	return ctx.Status(fiber.StatusUnauthorized).JSON(response)
}

func ServerError(ctx *fiber.Ctx) error {
	response := makeResponse(500, nil, "Internal server error")
	return ctx.Status(fiber.StatusInternalServerError).JSON(response)
}

func ServiceError(ctx *fiber.Ctx, name string, err error) error {
	var clientError *models.ClientErr
	if errors.As(err, &clientError) {
		return BadRequest(ctx, err)
	}
	logger.Debug("Error [", name, "]", err)
	return ServerError(ctx)
}

func ParsingError(ctx *fiber.Ctx, name string, parseErr error, validateErr error) error {
	logger.Warn("Error [", name, "]: Invalid post request from:", ctx.IP(), parseErr, validateErr)
	return BadRequest(ctx, errors.New("invalid item object"))
}

func makeResponse(status int, data any, message string) Response {
	return Response{
		Status:  status,
		Data:    data,
		Message: message,
	}
}
