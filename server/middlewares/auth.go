package middlewares

import (
	"server/token"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "missing access token"})
			context.Abort()
			return
		}
		claims, err := token.ValidateAccessToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Set("userID", claims.UserID)
		context.Next()
	}
}
