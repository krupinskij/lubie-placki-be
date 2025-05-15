package middlewares

import "github.com/gin-gonic/gin"

func Header() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	}
}
