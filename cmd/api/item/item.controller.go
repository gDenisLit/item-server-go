package item

import (
	"errors"

	"github.com/gDenisLit/item-server-go/cmd/models"

	"github.com/gDenisLit/item-server-go/cmd/services/logger"
	"github.com/gDenisLit/item-server-go/cmd/services/response"

	"github.com/gofiber/fiber/v2"
)

var clientError *models.ClientErr

func GetItems(ctx *fiber.Ctx) error {
	filterBy := models.FilterBy{
		Txt: ctx.Query("txt"),
	}
	items, err := itemService.query(filterBy)
	if err != nil {
		return response.ServerError(ctx)
	}
	return response.Success(ctx, items)
}

func GetItemById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return response.BadRequest(ctx, errors.New("invalid id"))
	}
	item, err := itemService.getById(id)
	if err != nil {
		return handleServiceError(ctx, "GetItemById", err)
	}
	return response.Success(ctx, item)
}

func RemoveItem(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return response.BadRequest(ctx, errors.New("invalid id"))
	}
	itemId, err := itemService.remove(id)
	if err != nil {
		return handleServiceError(ctx, "RemoveItem", err)
	}
	return response.Success(ctx, itemId)
}

func AddItem(ctx *fiber.Ctx) error {
	item := new(models.ItemDTO)
	parseErr := ctx.BodyParser(item)
	validateErr := item.Validate()

	if parseErr != nil || validateErr != nil {
		return handleParsingError(ctx, "AddItem", parseErr, validateErr)
	}
	savedItem, err := itemService.add(item)
	if err != nil {
		return handleServiceError(ctx, "AddItem", err)
	}
	return response.Success(ctx, savedItem)
}

func UpdateItem(ctx *fiber.Ctx) error {
	item := new(models.Item)
	parseErr := ctx.BodyParser(item)
	validateErr := item.Validate()

	if parseErr != nil || validateErr != nil {
		return handleParsingError(ctx, "UpdateItem", parseErr, validateErr)
	}
	savedItem, err := itemService.update(item)
	if err != nil {
		return handleServiceError(ctx, "UpdateItem", err)
	}
	return response.Success(ctx, savedItem)
}

func handleServiceError(ctx *fiber.Ctx, name string, err error) error {
	if errors.As(err, &clientError) {
		return response.BadRequest(ctx, err)
	}
	logger.Debug("Error [", name, "]", err)
	return response.ServerError(ctx)
}

func handleParsingError(ctx *fiber.Ctx, name string, parseErr error, validateErr error) error {
	logger.Warn("Error [", name, "]: Invalid post request from:", ctx.IP(), parseErr, validateErr)
	return response.BadRequest(ctx, errors.New("invalid item object"))
}
