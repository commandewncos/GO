package middlewares

import "github.com/gin-gonic/gin"

func AccessControlAllowOriginMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Next()
}
