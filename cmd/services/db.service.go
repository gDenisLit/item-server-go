package services

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database
var connected bool

func GetDBColletion(collection string) (*mongo.Collection, error) {
	if !connected {
		err := ConnectDB()
		if err != nil {
			return nil, err
		}
	}
	return db.Collection(collection), nil
}

func ConnectDB() error {
	if connected {
		return nil
	}

	dbUrl := os.Getenv("ATLAS_URL")
	dbName := os.Getenv("DB_NAME")

	if dbUrl == "" || dbName == "" {
		return errors.New("error connecting to DB: Missing dbUrl or dbName")
	}

	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(dbUrl),
	)
	if err != nil {
		return err
	}

	db = client.Database(dbName)
	connected = true

	// Log.Info("MongoDB Connection established")
	return nil
}
