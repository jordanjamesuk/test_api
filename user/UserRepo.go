package user

import (
	. "test_api/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepo interface {
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserRepo struct {
	Collection *mongo.Collection
}

func NewUserRepo(db *Db) *UserRepo {
	collection := db.Database.Collection("Users")

	return &UserRepo{
		Collection: collection,
	}
}

func (userRepo *UserRepo) GetUser(c *gin.Context) {
	id := c.Query("id")

	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"status": "failed", "message": id + " is not a valid objectID"})
		return
	}

	foundUser, err := FindUserByID(&objid, userRepo)
	if err != nil {
		c.JSON(404, gin.H{"status": "failed", "message": "Unable to find user: " + id})
		return
	}

	c.JSON(200, foundUser)
	return
}

func (userRepo *UserRepo) CreateUser(c *gin.Context) {
	var postBody User

	if err := c.BindJSON(&postBody); err != nil {
		c.JSON(400, gin.H{"status": "failed", "message": "Unable to validate incoming json body"})
		return
	}

	newUser, err := NewUser(postBody.Name, postBody.Email, userRepo)
	if err != nil {
		c.JSON(500, gin.H{"status": "failed", "message": "Unable to insert new user into database"})
		return
	}

	c.JSON(201, newUser)
}

func (userRepo *UserRepo) UpdateUser(c *gin.Context) {
	var postBody User

	if err := c.BindJSON(&postBody); err != nil {
		c.JSON(400, gin.H{"status": "failed", "message": "Unable to validate incoming json body"})
		return
	}

	foundUser, err := FindUserByID(&postBody.Id, userRepo)
	if err != nil {
		c.JSON(404, gin.H{"status": "failed", "message": "Unable to find user: " + postBody.Id.Hex()})
		return
	}

	hasDeletedUser, err := foundUser.UpdateUser(&postBody)
	if err != nil {
		c.JSON(500, gin.H{"status": "failed", "message": "Unable to update user: " + postBody.Id.Hex()})
		return
	}

	c.JSON(200, gin.H{"deleted": hasDeletedUser})
	return
}

func (userRepo *UserRepo) DeleteUser(c *gin.Context) {
	id := c.Query("id")

	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"status": "failed", "message": id + " is not a valid objectID"})
		return
	}

	foundUser, err := FindUserByID(&objid, userRepo)
	if err != nil {
		c.JSON(404, gin.H{"status": "failed", "message": "Unable to find user: " + id})
		return
	}

	hasDeletedUser, err := foundUser.DeleteUser()
	if err != nil {
		c.JSON(500, gin.H{"status": "failed", "message": "Unable to delete user: " + id})
		return
	}

	c.JSON(200, gin.H{"deleted": hasDeletedUser})
	return
}
