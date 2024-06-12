package routes

import (
	"server/controllers"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

func addPortfolioRoutes(rg *gin.RouterGroup) {

	portfolio := rg.Group("/portfolio")
	// Apply middleware to the /portfolio group
	portfolio.Use(middlewares.Auth())
	{
		portfolio.POST("/", controllers.CreatePortfolio)
		portfolio.GET("/:id", controllers.GetPortfolio)
		portfolio.PUT("/:id", controllers.UpdatePortfolio)
		portfolio.DELETE("/:id", controllers.DeletePortfolio)

		// positions := portfolio.Group("/:portfolio_id/positions")
		// {
		// 	positions.POST("/", controllers.CreatePosition)
		// 	positions.GET("/", controllers.GetPositions)
		// 	positions.GET("/:position_id", controllers.GetPositionByID)
		// 	positions.PUT("/:position_id", controllers.UpdatePosition)
		// 	positions.DELETE("/:position_id", controllers.DeletePosition)
		// }
	}

	addPositionsRoutes(portfolio)
}
