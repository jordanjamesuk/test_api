package server

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"os"
	. "test_api/database"
)

type Server struct {
	Database Db
	Router   *gin.Engine
}

func NewServer(database Db) *Server {
	sessionSecret := os.Getenv("SESSION_SECRET")
	router := gin.Default()

	server := &Server{
		Database: database,
		Router:   router,
	}

	// max age of cookie is 1 day
	store := cookie.NewStore([]byte(sessionSecret))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24})

	router.Use(sessions.Sessions("LoginSession", store))

	loginGroup := router.Group("/login")
	{
		loginGroup.POST("/", server.Login)
	}

	registerGroup := router.Group("/register")
	{
		registerGroup.POST("/", server.Register)
	}

	userGroup := router.Group("/user")
	{
		userGroup.Use(server.AuthRequired)
		userGroup.GET("/", server.UserHome)
	}

	return server
}
