package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetUserRoute(r *gin.Engine) {
	user := r.Group("/user")
	user.Use(middleware.AdminAuthMiddleware())
	{
		user.POST("/GetUserList", api.GetUserList)
		user.POST("/GetUser", api.GetUser)
		user.POST("/CreateUser", api.CreateUser)
		user.POST("/UpdateUser", api.UpdateUser)
		user.POST("/DeleteUser", api.DeleteUser)
	}
}
