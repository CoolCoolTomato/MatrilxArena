package middleware

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/common/response"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			response.Fail("", "Authorization header required", c)
			c.Abort()
			return
		}
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			response.Fail("", "Invalid token", c)
			c.Abort()
			return
		}
		c.Set("Username", claims.Username)
		c.Next()
	}
}
