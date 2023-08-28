package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"test_api/utils"
)

func (s *Server) Register(c *gin.Context) {
	var postBodyUser struct {
		Username string `json:"username" bson:"username" binding:"required" validate:"min=3,max=20"`
		Name     string `json:"name" bson:"name" binding:"required"`
		Email    string `json:"email" bson:"email" binding:"required" validate:"email"`
		Password string `json:"password" bson:"password" binding:"required"` // add some regex to make sure password is meeting requirements
	}

	if err := c.BindJSON(&postBodyUser); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"status": "failed", "message": "Unable to validate incoming json body"})
		return
	}

	validatorErr := validator.New().Struct(postBodyUser)
	if validatorErr != nil {
		c.JSON(400, gin.H{"status": "failed", "message": validatorErr.Error()})
		return
	}

	password := utils.HashAndSalt([]byte(postBodyUser.Password))

	s.Database.NewUser(postBodyUser.Username, postBodyUser.Name, postBodyUser.Email, password)
	return
}
