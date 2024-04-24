package routers

import (
	"go-bank/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	// users
	router.Group("/users")
	{
		router.GET("/", controllers.GetUsers)
		router.POST("/", controllers.CreateUser)
		router.PUT("/:id", controllers.UpdateUser)
		router.DELETE("/:id", controllers.DeleteUser)
	}

	return router
}
