package item

import (
	"errors"

	"github.com/gDenisLit/item-server-go/cmd/models"

	"github.com/gDenisLit/item-server-go/cmd/services/response"

	"github.com/gofiber/fiber/v2"
)

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
		return response.ServiceError(ctx, "GetItemById", err)
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
		return response.ServiceError(ctx, "RemoveItem", err)
	}
	return response.Success(ctx, itemId)
}

func AddItem(ctx *fiber.Ctx) error {
	item := new(models.ItemDTO)
	parseErr := ctx.BodyParser(item)
	validateErr := item.Validate()

	if parseErr != nil || validateErr != nil {
		return response.ParsingError(ctx, "AddItem", parseErr, validateErr)
	}
	savedItem, err := itemService.add(item)
	if err != nil {
		return response.ServiceError(ctx, "AddItem", err)
	}
	return response.Success(ctx, savedItem)
}

func UpdateItem(ctx *fiber.Ctx) error {
	item := new(models.Item)
	parseErr := ctx.BodyParser(item)
	validateErr := item.Validate()

	if parseErr != nil || validateErr != nil {
		return response.ParsingError(ctx, "UpdateItem", parseErr, validateErr)
	}
	savedItem, err := itemService.update(item)
	if err != nil {
		return response.ServiceError(ctx, "UpdateItem", err)
	}
	return response.Success(ctx, savedItem)
}
