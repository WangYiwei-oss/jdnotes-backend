package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CrossMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		if method != "" {
			context.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			context.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			context.Header("Access-Control-Allow-Credentials", "true")
			fmt.Println("设置了")
		}
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}
