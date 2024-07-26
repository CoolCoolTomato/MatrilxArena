package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetUserGroupContainerRoute(r *gin.Engine) {
	userContainer := r.Group("/userGroupContainer")
	userContainer.Use(middleware.JWTAuthMiddleware())
	{
		userContainer.POST("/GetGroupContainerListByUser", api.GetGroupContainerListByUser)
		userContainer.POST("/CreateGroupContainerByUser", api.CreateGroupContainerByUser)
		userContainer.POST("/DestroyGroupContainerByUser", api.DestroyGroupContainerByUser)
		userContainer.POST("/DelayGroupContainerByUser", api.DelayGroupContainerByUser)
	}
}
