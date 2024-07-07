package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetAuthRoute(r *gin.Engine) {
	noauth := r.Group("/auth")
	{
		noauth.POST("/login", api.Login)
	}
	auth := r.Group("/auth")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.POST("/GetUserByAuth", api.GetUserByAuth)
	}
}
