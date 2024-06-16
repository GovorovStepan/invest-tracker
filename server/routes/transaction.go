package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func addTransactionRoutes(rg *gin.RouterGroup) {

	positions := rg.Group("/:position_id/transaction")
	{
		positions.POST("/", controllers.CreateTransaction)
		positions.GET("/", controllers.GetTransactions)
		positions.PUT("/:transaction_id", controllers.UpdateTransaction)
		positions.DELETE("/:transaction_id", controllers.DeleteTransaction)
	}

}
