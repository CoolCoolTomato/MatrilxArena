package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetGroupRoute(r *gin.Engine) {
	adminGroup := r.Group("/group")
	adminGroup.Use(middleware.AdminAuthMiddleware())
	{
		adminGroup.POST("/GetGroupList", api.GetGroupList)
		adminGroup.POST("/GetGroup", api.GetGroup)
		adminGroup.POST("/CreateGroup", api.CreateGroup)
		adminGroup.POST("/UpdateGroup", api.UpdateGroup)
		adminGroup.POST("/DeleteGroup", api.DeleteGroup)
	}
}
