package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_project/gin_project/gin_test/response"
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
