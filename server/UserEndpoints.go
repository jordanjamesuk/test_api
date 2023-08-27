package server

import (
	. "test_api/user"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) GetUser(c *gin.Context) {
	id := c.Query("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"status": "failed", "message": id + " is not a valid objectID"})
		return
	}

	foundUser, err := s.Database.FindUserByID(&objID)
	if err != nil {
		c.JSON(404, gin.H{"status": "failed", "message": "Unable to find user: " + id})
		return
	}

	c.JSON(200, foundUser)
	return
}

func (s *Server) CreateUser(c *gin.Context) {
	var postBody User

	if err := c.BindJSON(&postBody); err != nil {
		c.JSON(400, gin.H{"status": "failed", "message": "Unable to validate incoming json body"})
		return
	}

	newUser, err := s.Database.NewUser(postBody.Name, postBody.Email)
	if err != nil {
		c.JSON(500, gin.H{"status": "failed", "message": "Unable to insert new user into database"})
		return
	}

	c.JSON(201, newUser)
}

func (s *Server) UpdateUser(c *gin.Context) {
	var postBody User

	if err := c.BindJSON(&postBody); err != nil {
		c.JSON(400, gin.H{"status": "failed", "message": "Unable to validate incoming json body"})
		return
	}

	updatedUser, err := s.Database.UpdateUser(&postBody.Id, &postBody)
	if err != nil {
		c.JSON(404, gin.H{"status": "failed", "message": "Unable to update user: " + postBody.Id.Hex()})
		return
	}

	c.JSON(200, updatedUser)
	return
}

func (s *Server) DeleteUser(c *gin.Context) {
	id := c.Query("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"status": "failed", "message": id + " is not a valid objectID"})
		return
	}

	hasDeletedUser, err := s.Database.DeleteUser(&objID)
	if err != nil {
		c.JSON(500, gin.H{"status": "failed", "message": "Unable to delete user: " + id})
		return
	}

	c.JSON(200, gin.H{"deleted": hasDeletedUser})
	return
}
