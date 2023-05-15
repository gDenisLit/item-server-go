package user

import (
	"context"

	"github.com/gDenisLit/item-server-go/cmd/models"
	"github.com/gDenisLit/item-server-go/cmd/services/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	collName string
}

var userService = &UserService{
	collName: "user",
}

func (s *UserService) query() ([]models.User, error) {

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

func (s *UserService) getById(userId string) (*models.User, error) {
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
		return nil, err
	}
	return user, nil
}

func (s *UserService) getByUsername(username string) (*models.User, error) {
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
		return nil, err
	}
	return user, nil
}

func (s *UserService) remove(id string) (*primitive.ObjectID, error) {
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

func (s *UserService) add(user *models.UserDTO) (*models.User, error) {
	collection, err := db.GetCollection(s.collName)
	if err != nil {
		return nil, err
	}
	res, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	objectId := res.InsertedID.(primitive.ObjectID)
	savedUser := &models.User{
		ID:       objectId,
		Username: user.Username,
		Password: user.Password,
		Fullname: user.Fullname,
		ImgUrl:   user.ImgUrl,
		IsAdmin:  false,
	}
	return savedUser, nil
}
