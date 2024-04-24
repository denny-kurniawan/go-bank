package routers

import (
	"go-bank/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	users := router.Group("/users")
	{
		users.POST("/register", controllers.Register)
		users.POST("/login", controllers.Login)
		users.POST("change-password", controllers.ChangePassword)
		users.DELETE("/", controllers.DeleteUser)
	}

	return router
}
