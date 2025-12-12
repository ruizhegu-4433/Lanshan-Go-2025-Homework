package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// 自定义中间件，记录请求的处理时间
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求开始时间
		start := time.Now()

		// 请求处理前的操作
		log.Println("请求开始")

		// 执行请求处理
		c.Next()

		// 请求处理后的操作
		duration := time.Since(start)
		log.Printf("请求 %s %s 花费了 %v\n", c.Request.Method, c.Request.URL.Path, duration)
	}
}
