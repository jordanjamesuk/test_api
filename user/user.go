package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `json:"_id" bson:"_id"`
	Username     string             `json:"username" bson:"username" binding:"required"`
	Name         string             `json:"name" bson:"name" binding:"required"`
	Email        string             `json:"email" bson:"email" binding:"required"`
	PasswordHash string             `json:"password_hash" bson:"password_hash" binding:"required"`
}
