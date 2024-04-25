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

	accounts := router.Group("/accounts")
	{
		accounts.POST("/", controllers.CreateAccount)

		userAccounts := accounts.Group("/:userID")
		{
			userAccounts.GET("/", controllers.GetAccountsByUserID)
			userAccounts.DELETE("/", controllers.DeleteAccountsByUserID)

			userAccount := userAccounts.Group("/:accountNo")
			{
				userAccount.GET("/", controllers.GetAccountDetails)
				userAccount.DELETE("/", controllers.DeleteAccountByAccountNo)
			}
		}
	}

	transactions := router.Group("/transactions")
	{
		transactions.POST("/deposit", controllers.Deposit)
		transactions.POST("/withdraw", controllers.Withdraw)
		transactions.POST("/transfer", controllers.Transfer)
	}

	return router
}
