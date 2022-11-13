package middleware

import (
	"fmt"
	"gin_project/response"
	"github.com/gin-gonic/gin"
)

func TokenMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Token")
		if token != "" {
			context.Next()
		}
		fmt.Println("没有token")
		response.FailAndMsg("没有token", context)
		context.Abort()
	}
}
