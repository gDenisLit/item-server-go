package dtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserDTO struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Fullname string             `json:"fullname" bson:"fullname"`
	ImgUrl   string             `json:"imgUrl" bson:"imgUrl"`
}
