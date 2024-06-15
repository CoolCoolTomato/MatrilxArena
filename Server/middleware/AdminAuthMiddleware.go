package middleware

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/jwt"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
	"strings"
)

func AdminAuthMiddleware() gin.HandlerFunc {
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

		var user model.User
		user.Username = claims.Username
		err = user.GetUserByUsernameOrEmail()
		if err != nil {
			response.Fail("", "Get user fail", c)
			c.Abort()
			return
		}

		if user.Role != 1 {
			response.Fail("", "Permission denied", c)
			c.Abort()
			return
		}

		c.Set("Username", user.Username)
		c.Next()
	}
}
