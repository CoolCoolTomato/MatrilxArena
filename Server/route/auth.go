package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/gin-gonic/gin"
)

func SetAuthRoute(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", api.Login)
	}
}
