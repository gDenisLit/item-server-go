package user

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

var UserService = &service{
	collName: "user",
}

func (s *service) Query() ([]models.User, error) {
	collection, err := db.GetCollection(s.collName)
	if err != nil {
		return nil, err
	}
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	var users []models.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (s *service) GetById(userId string) (*models.User, error) {
	collection, err := db.GetCollection(s.collName)
	if err != nil {
		return nil, err
	}
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, &models.ClientErr{Message: "invalid id"}
	}
	user := &models.User{}
	err = collection.FindOne(
		context.TODO(),
		bson.M{"_id": objectId},
	).Decode(user)

	if err != nil {
		return nil, &models.ClientErr{Message: "invalid id"}
	}
	return user, nil
}

func (s *service) GetByUsername(username string) (*models.User, error) {
	collection, err := db.GetCollection(s.collName)
	if err != nil {
		return nil, err
	}
	user := &models.User{}
	err = collection.FindOne(
		context.TODO(),
		bson.M{"username": username},
	).Decode(user)

	if err != nil {
		return nil, &models.ClientErr{Message: "invalid username"}
	}
	return user, nil
}

func (s *service) Remove(id string) (*primitive.ObjectID, error) {
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

func (s *service) Add(user *models.User) (*models.User, error) {
	collection, err := db.GetCollection(s.collName)
	if err != nil {
		return nil, err
	}
	res, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	objectId := res.InsertedID.(primitive.ObjectID)
	user.ID = objectId
	user.IsAdmin = false
	return user, nil
}
