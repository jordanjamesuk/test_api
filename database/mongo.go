package database

import (
	"context"
	. "test_api/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewMongoDatabase(connectionString string) (*MongoDB, error) {
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

	return &MongoDB{
		Client:   client,
		Database: client.Database("test_api"),
	}, nil

}

func (db *MongoDB) NewUser(name, email string) (newUser *User, err error) {
	result, err := db.Database.Collection("users").InsertOne(context.Background(), bson.M{"name": name, "email": email})
	if err != nil {
		return nil, err
	}

	return &User{
		Id:    result.InsertedID.(primitive.ObjectID),
		Name:  name,
		Email: email,
	}, nil
}

func (db *MongoDB) FindUserByID(id *primitive.ObjectID) (foundUser *User, err error) {
	var result User

	if err = db.Database.Collection("users").FindOne(context.Background(), bson.M{"_id": id}).Decode(&result); err == mongo.ErrNoDocuments {
		return nil, err
	}

	return &result, nil
}

func (db *MongoDB) UpdateUser(id *primitive.ObjectID, updatedUser *User) (userUpdated *User, err error) {
	updatedUser.Id = *id

	var result User
	err = db.Database.Collection("users").FindOneAndUpdate(context.Background(), bson.M{"_id": id}, bson.M{"$set": updatedUser}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (db *MongoDB) DeleteUser(id *primitive.ObjectID) (isSuccessful bool, err error) {
	_, err = db.Database.Collection("users").DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return false, err
	}

	return true, nil
}
