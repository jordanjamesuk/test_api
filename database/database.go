package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	. "test_api/user"
)

type Db interface {
	NewUser(username, name, email string, passwordHash string) (newUser *User, err error)
	GetAllUsers() (users *[]User, err error)
	FindUserByKeyValue(key string, value any) (foundUser *User, err error)
	UpdateUser(id *primitive.ObjectID, updatedUser *User) (userUpdated *User, err error)
	DeleteUser(id *primitive.ObjectID) (isSuccessful bool, err error)
}
