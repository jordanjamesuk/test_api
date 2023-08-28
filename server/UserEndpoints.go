package server

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) UserHome(c *gin.Context) {
	session := sessions.Default(c)
	userIDHex := session.Get("user")
	userID, err := primitive.ObjectIDFromHex(userIDHex.(string))
	if err != nil {
		c.JSON(500, gin.H{"status": "failed", "message": "Might of deleted user, while session still exists"})
		return
	}

	user, err := s.Database.FindUserByKeyValue("_id", userID)
	if err != nil {
		c.JSON(500, gin.H{"status": "failed", "message": "Might of deleted user, while session still exists"})
		return
	}

	c.JSON(200, gin.H{"user": user})
}

