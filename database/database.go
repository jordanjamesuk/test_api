package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Db struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewDatabase(connectionString string) (*Db, error) {
	println("[INFO]: Attempting to connect to database...")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))

	if err != nil {
		println("[ERR]: Invalid Mongo URL")
		return nil, err
	}

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		println("[ERR]: Could not ping database... Is the database running?")
		return nil, err
	}

	println("[INFO]: Connected to Database")

	return &Db{
		Client:   client,
		Database: client.Database("test_api"),
	}, nil
}
