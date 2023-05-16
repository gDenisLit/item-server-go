package item

import (
	"context"

	"github.com/gDenisLit/item-server-go/cmd/models"
	"github.com/gDenisLit/item-server-go/cmd/services/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	collName string
}

var ItemService = &service{
	collName: "item",
}

func (s *service) query(filterBy models.FilterBy) ([]models.Item, error) {
	criteria := buildCriteria(filterBy)
	collection, err := db.GetCollection(s.collName)
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

func (s *service) getById(itemId string) (*models.Item, error) {
	collection, err := db.GetCollection(s.collName)
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

func (s *service) remove(id string) (*primitive.ObjectID, error) {
	collection, err := db.GetCollection(s.collName)
	if err != nil {
		return nil, err
	}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, &models.ClientErr{Message: "invalid id"}
	}
	res, err := collection.DeleteOne(
		context.TODO(),
		bson.M{"_id": objectId},
	)
	if err != nil || res.DeletedCount == 0 {
		return nil, &models.ClientErr{Message: "invalid id"}
	}
	return &objectId, nil
}

func (s *service) add(item *models.ItemDTO) (*models.Item, error) {
	collection, err := db.GetCollection(s.collName)
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

func (s *service) update(item *models.Item) (*models.Item, error) {
	collection, err := db.GetCollection(s.collName)
	if err != nil {
		return nil, err
	}
	res, err := collection.ReplaceOne(
		context.TODO(),
		bson.M{"_id": item.ID},
		item,
	)
	if err != nil || res.UpsertedCount == 0 {
		return nil, &models.ClientErr{Message: "invalid item object"}
	}
	return item, nil
}
