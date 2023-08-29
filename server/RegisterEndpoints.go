package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"test_api/utils"
)

func (s *Server) Register(c *gin.Context) {
	if c.ContentType() != "x-www-form-urlencoded" {
		c.JSON(415, gin.H{"status": "failed", "message": "Only accept content ContentType of x-www-form-urlencoded"})
		return
	}

	var postBodyUser struct {
		Username string `form:"username" binding:"required" validate:"min=3,max=20"`
		Name     string `form:"name" binding:"required"`
		Email    string `form:"email" binding:"required" validate:"email"`
		Password string `form:"password" binding:"required"` // add some regex to make sure password is meeting requirements
	}

	if err := c.ShouldBind(&postBodyUser); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"status": "failed", "message": "Unable to validate incoming body"})
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
