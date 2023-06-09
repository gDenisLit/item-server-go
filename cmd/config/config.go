package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DB_URL      string
	DB_NAME     string
	SECRET_KEY  string
	BLOCK_KEY   string
	DEV_ENV     string
	SALT_ROUNDS int
	PORT        string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env")
	}
	DB_URL = os.Getenv("ATLAS_URL")
	DB_NAME = os.Getenv("DB_NAME")
	SECRET_KEY = os.Getenv("CRYPTER_KEY")
	BLOCK_KEY = os.Getenv("BLOCK_KEY")
	DEV_ENV = os.Getenv("DEV_ENV")
	SALT_ROUNDS, err = strconv.Atoi(os.Getenv("SALT_ROUNDS"))
	if err != nil {
		log.Fatal("Failed to load env")
	}
	PORT = os.Getenv("PORT")
}
