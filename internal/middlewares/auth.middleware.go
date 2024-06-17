package middlewares

import (
	"github.com/KiAnh2911/go-crm-api-shop/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token != "vaild-token" {
			response.ErrorRespone(c, response.ErrorInvaildToken, "")
			c.Abort()
			return
		}
		c.Next()
	}
}
