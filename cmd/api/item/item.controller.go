package item

import (
	"errors"

	"github.com/gDenisLit/item-server-go/cmd/dtos"
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
		if errors.As(err, &clientError) {
			return response.BadRequest(ctx, err)
		}
		logger.Debug("Error [GetItemById]:", err, id)
		return response.ServerError(ctx)
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
		if errors.As(err, &clientError) {
			return response.BadRequest(ctx, err)
		}
		logger.Debug("Error [RemoveItem]:", err, id)
		return response.ServerError(ctx)
	}
	return response.Success(ctx, itemId)
}

func AddItem(ctx *fiber.Ctx) error {
	item := new(models.ItemDTO)
	parseErr := ctx.BodyParser(item)
	validateErr := item.Validate()

	if parseErr != nil || validateErr != nil {
		logger.Warn("Error [AddItem]: Invalid post request", parseErr, validateErr)
		return response.BadRequest(ctx, errors.New("invalid item object"))
	}
	savedItem, err := itemService.add(item)
	if err != nil {
		if errors.As(err, &clientError) {
			return response.BadRequest(ctx, err)
		}
		logger.Debug("Error [AddItem]:", err)
		return response.ServerError(ctx)
	}
	return response.Success(ctx, savedItem)
}

func UpdateItem(ctx *fiber.Ctx) error {
	item := new(dtos.UpdateItemDTO)
	err := ctx.BodyParser(item)

	if err != nil {
		logger.Warn("Error [UpdateItem]: Invalid post request", err)
		return response.BadRequest(ctx, errors.New("invalid item object"))
	}
	savedItem, err := itemService.update(item)
	if err != nil {
		if errors.As(err, &clientError) {
			return response.BadRequest(ctx, err)
		}
		logger.Debug("Error [UpdateItem]:", err)
		return response.ServerError(ctx)
	}
	return response.Success(ctx, savedItem)
}
