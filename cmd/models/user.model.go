package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Fullname string             `json:"fullname" bson:"fullname"`
	ImgUrl   string             `json:"imgUrl" bson:"imgUrl"`
	IsAdmin  bool               `json:"isAdmin" bson:"isAdmin"`
}
