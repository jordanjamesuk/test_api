package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name" bson:"name" binding:"required"`
	Email    string             `json:"email" bson:"email" binding:"required"`
	UserRepo *UserRepo          `json:"-"`
}

func NewUser(name, email string, userRepo *UserRepo) (newUser *User, err error) {
	result, err := userRepo.Collection.InsertOne(context.Background(), bson.M{"name": name, "email": email})
	if err != nil {
		return nil, err
	}

	return &User{
		Id:       result.InsertedID.(primitive.ObjectID),
		Name:     name,
		Email:    email,
		UserRepo: userRepo,
	}, nil
}

func FindUserByID(id *primitive.ObjectID, userRepo *UserRepo) (foundUser *User, err error) {
	var result User

	if err = userRepo.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&result); err == mongo.ErrNoDocuments {
		return nil, err
	}

	result.UserRepo = userRepo

	return &result, nil
}

func (user *User) UpdateUser(updatedUser *User) (modifiedUsersCount int64, err error) {
	updatedUser.Id = user.Id // probably a better way to do this? this is so it doesnt complain about changing the _id, as updatedUser.Id is blank

	result, err := user.UserRepo.Collection.UpdateOne(context.Background(), bson.M{"_id": user.Id}, bson.M{"$set": updatedUser})
	if err != nil {
		return 0, err
	}

	return result.MatchedCount, nil
}

func (user *User) DeleteUser() (isSuccessful bool, err error) {
	_, err = user.UserRepo.Collection.DeleteOne(context.Background(), bson.M{"_id": user.Id})
	if err != nil {
		return false, err
	}

	return true, nil
}
