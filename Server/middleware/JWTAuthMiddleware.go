package middleware

import (
    "github.com/CoolCoolTomato/MatrilxArena/Server/utils/jwt"
    "github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
    "github.com/gin-gonic/gin"
    "strings"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			response.Fail("", "Authorization header required", c)
			c.Abort()
			return
		}

		if strings.HasPrefix(tokenString, "Bearer ") {
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		}

		claims, err := jwt.ValidateJWT(tokenString)
		if err != nil {
			response.Fail("", "Invalid token", c)
			c.Abort()
			return
		}
		c.Set("Username", claims.Username)
		c.Next()
	}
}
