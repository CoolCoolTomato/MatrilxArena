package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetUserGroupRoute(r *gin.Engine) {
	userGroup := r.Group("/userGroup")
	userGroup.Use(middleware.JWTAuthMiddleware())
	{
		userGroup.POST("/GetUserGroupList", api.GetUserGroupList)
		userGroup.POST("/GetVisibleUserGroupList", api.GetVisibleUserGroupList)
		userGroup.POST("/AddUserGroup", api.AddUserGroup)
		userGroup.POST("/RemoveUserGroup", api.RemoveUserGroup)
	}
}
