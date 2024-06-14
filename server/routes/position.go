package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func addPositionsRoutes(rg *gin.RouterGroup) {

	positions := rg.Group("/:portfolio_id/position")
	{
		positions.POST("/", controllers.CreatePosition)
		positions.GET("/", controllers.GetPositions)
		positions.GET("/:position_id", controllers.GetPositionByID)
		positions.PUT("/:position_id", controllers.UpdatePosition)
		positions.DELETE("/:position_id", controllers.DeletePosition)
	}

}
