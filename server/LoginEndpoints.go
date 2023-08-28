package server

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	. "test_api/user"
	"test_api/utils"
)

func (s *Server) Login(c *gin.Context) {
	session := sessions.Default(c)

	var postBodyUser struct {
		Username string `json:"username" bson:"username" validate:"required_without=Email"`
		Email    string `json:"email" bson:"email" validate:"required_without=Username, email"`
		Password string `json:"password" bson:"password" binding:"required"`
	}

	if err := c.BindJSON(&postBodyUser); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"status": "failed", "message": "Unable to validate incoming json body"})
		return
	}

	validateErr := validator.New().Struct(postBodyUser)
	if validateErr != nil {
		fmt.Println(validateErr)
		c.JSON(400, gin.H{"status": "failed", "message": "Unable to validate incoming json body"})
		return
	}

	var user *User
	var err error

	if postBodyUser.Username != "" {
		user, err = s.Database.FindUserByKeyValue("username", postBodyUser.Username)
		if err != nil {
			c.JSON(400, gin.H{"status": "failed", "message": "Unable to find user"})
			return
		}
	} else {
		user, err = s.Database.FindUserByKeyValue("Email", postBodyUser.Email)
		if err != nil {
			c.JSON(400, gin.H{"status": "failed", "message": "Unable to find user"})
			return
		}
	}

	if success := utils.ComparePasswords(user.PasswordHash, []byte(postBodyUser.Password)); success == false {
		c.JSON(401, gin.H{"status": "failed", "message": "Incorrect password"})
		return
	}

	session.Set("user", user.Id.Hex())
	if err := session.Save(); err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"status": "failed", "message": "Failed to save session"})
		return
	}

	c.JSON(200, gin.H{"status": "success", "message": "Successfully logged user in"})
	return
}
