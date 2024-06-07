package middlewares

import (
	"github.com/gin-gonic/gin"

	"server/token"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "missing access token"})
			context.Abort()
			return
		}
		err := token.ValidateAccessToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
