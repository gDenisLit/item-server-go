package item

import (
	"context"

	"github.com/gDenisLit/item-server-go/cmd/dtos"
	"github.com/gDenisLit/item-server-go/cmd/models"
	"github.com/gDenisLit/item-server-go/cmd/services/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const collName string = "item"

type ItemService struct{}

var itemService = &ItemService{}

func (s *ItemService) query(filterBy models.FilterBy) ([]models.Item, error) {

	criteria := buildCriteria(filterBy)
	collection, err := db.GetCollection(collName)
	if err != nil {
		return nil, err
	}

	cursor, err := collection.Find(context.TODO(), criteria)
	if err != nil {
		return nil, err
	}

	var items []models.Item
	if err = cursor.All(context.TODO(), &items); err != nil {
		return nil, err
	}
	return items, nil
}

func buildCriteria(filterBy models.FilterBy) bson.M {
	criteria := bson.M{
		"name": primitive.Regex{Pattern: filterBy.Txt, Options: "i"},
	}
	return criteria
}

func (s *ItemService) getById(itemId string) (*models.Item, error) {
	collection, err := db.GetCollection(collName)
	if err != nil {
		return nil, err
	}

	objectId, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return nil, &models.ClientErr{Message: "invalid id"}
	}

	item := &models.Item{}
	err = collection.FindOne(
		context.TODO(),
		bson.M{"_id": objectId},
	).Decode(item)

	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *ItemService) remove(id string) (*primitive.ObjectID, error) {
	collection, err := db.GetCollection(collName)
	if err != nil {
		return nil, err
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, &models.ClientErr{Message: "invalid id"}
	}

	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		return nil, err
	}
	return &objectId, nil
}

func (s *ItemService) add(item *models.ItemDTO) (*models.Item, error) {
	collection, err := db.GetCollection(collName)
	if err != nil {
		return nil, err
	}

	res, err := collection.InsertOne(context.TODO(), item)
	if err != nil {
		return nil, err
	}

	objectId := res.InsertedID.(primitive.ObjectID)
	savedItem := &models.Item{
		ID:     objectId,
		Name:   item.Name,
		Price:  item.Price,
		ImgUrl: item.ImgUrl,
	}
	return savedItem, nil
}

func (s *ItemService) update(item *dtos.UpdateItemDTO) (*models.Item, error) {
	collection, err := db.GetCollection(collName)
	if err != nil {
		return nil, err
	}

	objectId, err := primitive.ObjectIDFromHex(item.ID)
	if err != nil {
		return nil, &models.ClientErr{Message: "invalid id"}
	}

	savedItem := &models.Item{
		ID:     objectId,
		Name:   item.Name,
		Price:  item.Price,
		ImgUrl: item.ImgUrl,
	}

	_, err = collection.ReplaceOne(
		context.TODO(),
		bson.M{"_id": objectId},
		savedItem,
	)
	if err != nil {
		return nil, err
	}
	return savedItem, nil
}
