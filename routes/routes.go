package routes

import (
	"meu-primeiro-crude-go/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/users", handlers.GetUsers)
	r.POST("/users", handlers.CreateUser)

	return r
}
