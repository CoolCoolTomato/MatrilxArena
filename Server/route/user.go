package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/gin-gonic/gin"
)

func SetUserRoute(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/GetUser", api.GetUser)
		user.POST("/CreateUser", api.CreateUser)
		user.POST("/UpdateUser", api.UpdateUser)
		user.POST("/DeleteUser", api.DeleteUser)
	}
}
