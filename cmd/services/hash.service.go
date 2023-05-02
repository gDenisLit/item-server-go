package services

import (
	"os"

	"github.com/gorilla/securecookie"
)

var (
	hashKey  []byte
	blockKey []byte
	secure   *securecookie.SecureCookie
)

func Encode(name string, value interface{}) (string, error) {
	s := getSecure()
	return s.Encode(name, value)
}

func Decode(name string, value string, dst interface{}) error {
	s := getSecure()
	return s.Decode(name, value, dst)
}

func getSecure() *securecookie.SecureCookie {
	if len(hashKey) == 0 && len(blockKey) == 0 {
		hashKey = []byte(os.Getenv("CRYPTER_KEY"))
		blockKey = []byte(os.Getenv("BLOCK_KEY"))
		secure = securecookie.New(hashKey, blockKey)
	}
	return secure
}
