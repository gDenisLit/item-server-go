package auth

import (
	"encoding/json"

	"github.com/gDenisLit/item-server-go/cmd/api/user"
	"github.com/gDenisLit/item-server-go/cmd/config"
	"github.com/gDenisLit/item-server-go/cmd/models"
	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	hashKey   []byte
	blockKey  []byte
	tokenName string
}

var AuthService = &service{
	hashKey:   []byte(config.SECRET_KEY),
	blockKey:  []byte(config.BLOCK_KEY),
	tokenName: "loginToken",
}

var secure = securecookie.New(
	AuthService.hashKey,
	AuthService.blockKey,
)

func (s *service) Login(credentials *models.User) (*models.User, error) {
	user, err := user.UserService.GetByUsername(credentials.Username)
	if err != nil {
		return nil, &models.ClientErr{Message: "invalid username or password"}
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(credentials.Password),
	)
	if err != nil {
		return nil, &models.ClientErr{Message: "invalid username or password"}
	}
	return user, nil
}

func (s *service) Signup(credentials *models.User) (*models.User, error) {
	_, err := user.UserService.GetByUsername(credentials.Username)
	if err == nil {
		return nil, &models.ClientErr{Message: "username is already taken"}
	}
	saltRounds := config.SALT_ROUNDS
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(credentials.Password),
		saltRounds,
	)
	if err != nil {
		return nil, err
	}
	credentials.Password = string(hash)
	user, err := user.UserService.Add(credentials)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) GetLoginToken(user *models.User) (string, error) {
	userJson, err := json.Marshal(user.Minify())
	if err != nil {
		return "", err
	}
	token, err := secure.Encode(s.tokenName, userJson)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *service) ValidateToken(token string) (*models.User, error) {
	var userBytes []byte
	err := secure.Decode(s.tokenName, token, &userBytes)
	if err != nil {
		return nil, err
	}
	var user models.User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
