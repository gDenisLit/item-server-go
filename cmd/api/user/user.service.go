package user

import (
	"context"

	"github.com/gDenisLit/item-server-go/cmd/dtos"
	"github.com/gDenisLit/item-server-go/cmd/models"
	"github.com/gDenisLit/item-server-go/cmd/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const collName string = "user"

func Query() ([]models.User, error) {

	collection, err := services.GetDBColletion(collName)
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

func GetById(userId string) (*models.User, error) {
	collection, err := services.GetDBColletion(collName)
	if err != nil {
		return nil, err
	}

	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	user := &models.User{}
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetByUsername(username string) (*models.User, error) {
	collection, err := services.GetDBColletion(collName)
	if err != nil {
		return nil, err
	}

	user := &models.User{}
	err = collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func Add(user *dtos.SignupDTO) (*models.User, error) {
	collection, err := services.GetDBColletion(collName)
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
		IsAdmin:  user.IsAdmin,
	}
	return savedUser, nil
}
