package models

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	Name   string             `json:"name" bson:"name"`
	Price  int                `json:"price" bson:"price"`
	ImgUrl string             `json:"imgUrl" bson:"imgUrl"`
}

type ItemDTO struct {
	Name   string `json:"name" bson:"name"`
	Price  int    `json:"price" bson:"price"`
	ImgUrl string `json:"imgUrl" bson:"imgUrl"`
}

func (item *ItemDTO) Validate() error {
	err := errors.New("missing required propery")

	if item.Name == "" {
		return err
	}
	if item.Price == 0 {
		return err
	}
	if item.ImgUrl == "" {
		return err
	}
	return nil
}

func (item *Item) Validate() error {
	err := errors.New("missing required propery")
	if item.ID == primitive.NilObjectID {
		return err
	}
	if item.Name == "" {
		return err
	}
	if item.Price == 0 {
		return err
	}
	if item.ImgUrl == "" {
		return err
	}
	return nil
}
