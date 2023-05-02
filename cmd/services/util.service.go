package services

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	// check if prod
	prod := os.Getenv("PROD")

	// if not, load .env
	if prod != "true" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}
