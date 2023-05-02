package auth

import (
	"encoding/json"
	"errors"

	"github.com/gDenisLit/item-server-go/cmd/api/user"
	"github.com/gDenisLit/item-server-go/cmd/dtos"
	"github.com/gDenisLit/item-server-go/cmd/models"
	"github.com/gDenisLit/item-server-go/cmd/services"
	"golang.org/x/crypto/bcrypt"
)

func Login(credentials *dtos.LoginDTO) (*models.User, error) {
	user, err := user.GetByUsername(credentials.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(credentials.Password),
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func Signup(credentials *dtos.SignupDTO) (*models.User, error) {
	_, err := user.GetByUsername(credentials.Username)
	if err == nil {
		return nil, errors.New("username already taken")
	}

	saltRounds := 10
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(credentials.Password),
		saltRounds,
	)
	if err != nil {
		return nil, err
	}

	credentials.Password = string(hash)
	user, err := user.Add(credentials)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetLoginToken(user *dtos.UserDTO) (string, error) {

	userJson, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	token, err := services.Encode("loginToken", userJson)
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidateToken(token string) (*models.User, error) {

	return nil, nil
}
