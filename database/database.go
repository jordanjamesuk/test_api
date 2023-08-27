package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	. "test_api/user"
)

type Db interface {
	NewUser(name, email string) (newUser *User, err error)
	FindUserByID(id *primitive.ObjectID) (foundUser *User, err error)
	UpdateUser(id *primitive.ObjectID, updatedUser *User) (userUpdated *User, err error)
	DeleteUser(id *primitive.ObjectID) (isSuccessful bool, err error)
}
