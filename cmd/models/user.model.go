package models

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Fullname string             `json:"fullname" bson:"fullname"`
	ImgUrl   string             `json:"imgUrl" bson:"imgUrl"`
	IsAdmin  bool               `json:"isAdmin" bson:"isAdmin"`
}

type UserDTO struct {
	Username string `json:"username" bson:"username"`
	Fullname string `json:"fullname" bson:"fullname"`
	ImgUrl   string `json:"imgUrl" bson:"imgUrl"`
}

func (user *User) ValidateLoginCredentials() error {
	err := errors.New("missing username or passowrd")
	if user.Username == "" {
		return err
	}
	if user.Password == "" {
		return err
	}
	return nil
}

func (user *User) ValidateSignupCredentials() error {
	err := errors.New("missing required propery")
	if user.Username == "" {
		return err
	}
	if user.Password == "" {
		return err
	}
	if user.Fullname == "" {
		return err
	}
	if user.ImgUrl == "" {
		return err
	}
	return nil
}

func (user *User) Validate() error {
	err := errors.New("missing required propery")
	if user.ID == primitive.NilObjectID {
		return err
	}
	if user.Username == "" {
		return err
	}
	if user.Password == "" {
		return err
	}
	if user.Fullname == "" {
		return err
	}
	if user.ImgUrl == "" {
		return err
	}
	return nil
}

func (user *UserDTO) Validate() error {
	err := errors.New("missing required propery")
	if user.Username == "" {
		return err
	}
	if user.Fullname == "" {
		return err
	}
	if user.ImgUrl == "" {
		return err
	}
	return nil
}

func (user *User) Minify() map[string]interface{} {
	minifiedUser := map[string]interface{}{
		"_id":      user.ID,
		"username": user.Username,
		"fullname": user.Fullname,
		"imgUrl":   user.ImgUrl,
	}
	return minifiedUser
}
