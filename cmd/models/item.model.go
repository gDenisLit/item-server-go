package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	Name   string             `json:"name" bson:"name"`
	Price  int                `json:"price" bson:"price"`
	ImgUrl string             `json:"imgUrl" bson:"imgUrl"`
}
