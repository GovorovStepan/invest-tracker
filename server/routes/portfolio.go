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
		portfolio.GET("/:portfolio_id", controllers.GetPortfolio)
		portfolio.PUT("/:portfolio_id", controllers.UpdatePortfolio)
		portfolio.DELETE("/:portfolio_id", controllers.DeletePortfolio)
	}

	addPositionsRoutes(portfolio)
}
