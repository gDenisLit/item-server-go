package item

import (
	"github.com/gDenisLit/item-server-go/cmd/dtos"
	"github.com/gDenisLit/item-server-go/cmd/models"
	"github.com/gDenisLit/item-server-go/cmd/services"
	"github.com/gofiber/fiber/v2"
)

func GetItems(ctx *fiber.Ctx) error {

	filterBy := models.FilterBy{
		Txt: ctx.Query("txt"),
	}

	services.Log.Info("Getting items", filterBy)
	channel := make(chan []models.Item)

	go func() {
		items, err := Query(filterBy)
		if err != nil {
			services.Log.Error("Item controller Error:", err.Error())
			channel <- nil
		} else {
			channel <- items
		}
	}()

	items := <-channel
	if items == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(items)
}

func GetItemById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	services.Log.Info("Getting item with id:", id)
	item, err := GetById(id)

	if err != nil {
		services.Log.Error("Item controller Error:", err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(item)
}

func AddItem(ctx *fiber.Ctx) error {
	item := new(dtos.AddItemDTO)
	err := ctx.BodyParser(item)

	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": "invalid item object",
		})
	}

	services.Log.Info("Adding new item", item)
	savedItem, err := Add(item)

	if err != nil {
		services.Log.Error("Item controller Error:", err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return ctx.Status(fiber.StatusAccepted).JSON(savedItem)
}

func UpdateItem(ctx *fiber.Ctx) error {

	item := new(dtos.UpdateItemDTO)
	err := ctx.BodyParser(item)

	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": "invalid item object",
		})
	}

	services.Log.Info("Updating item", item)
	savedItem, err := Update(item)
	if err != nil {
		services.Log.Error("Item controller Error:", err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return ctx.Status(fiber.StatusAccepted).JSON(savedItem)
}

func RemoveItem(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	services.Log.Info("Removing item with id:", id)
	itemId, err := Remove(id)

	if err != nil {
		services.Log.Error("Item controller Error:", err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return ctx.Status(fiber.StatusAccepted).JSON(itemId)
}
