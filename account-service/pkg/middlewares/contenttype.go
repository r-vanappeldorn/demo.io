package middlewares

import "github.com/gin-gonic/gin"

func ContentType(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Next()
}
