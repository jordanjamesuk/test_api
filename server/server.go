package server

import (
	"github.com/gin-gonic/gin"
	. "test_api/database"
	. "test_api/user"
)

type Server struct {
	UserRepo *UserRepo
	Database *Db
	Router   *gin.Engine
}

func NewServer(database *Db) *Server {
	userRepo := NewUserRepo(database)
	router := gin.Default()

	server := &Server{
		UserRepo: userRepo,
		Database: database,
		Router:   router,
	}

	userGroup := router.Group("/user")
	{
		userGroup.GET("/", userRepo.GetUser)
		userGroup.POST("/", userRepo.CreateUser)
		userGroup.PUT("/", userRepo.UpdateUser)
		userGroup.DELETE("/", userRepo.DeleteUser)
	}

	return server
}
