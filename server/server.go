package server

import (
	"github.com/gin-gonic/gin"
	. "test_api/database"
)

type Server struct {
	Database Db
	Router   *gin.Engine
}

func NewServer(database Db) *Server {
	router := gin.Default()

	server := &Server{
		Database: database,
		Router:   router,
	}

	userGroup := router.Group("/user")
	{
		userGroup.GET("/", server.GetUser)
		userGroup.POST("/", server.CreateUser)
		userGroup.PUT("/", server.UpdateUser)
		userGroup.DELETE("/", server.DeleteUser)
	}

	return server
}
