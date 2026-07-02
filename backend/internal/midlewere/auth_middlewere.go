package midlewere

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckingAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error" : "needed authorization token",
			})
			return 
		}

		if !strings.HasPrefix(token, "Bearer ") {
			ctx.JSON(401, gin.H{
				"error" : "invalid authorization header",
			})
		}
		ctx.Next()
	}
}