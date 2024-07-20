package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetGroupRoute(r *gin.Engine) {
	group := r.Group("/group")
    group.Use(middleware.JWTAuthMiddleware())
	{
		group.POST("/GetGroupList", api.GetGroupList)
		group.POST("/GetGroup", api.GetGroup)
	}
    adminGroup := r.Group("/group")
    adminGroup.Use(middleware.AdminAuthMiddleware())
    {
        adminGroup.POST("/CreateGroup", api.CreateGroup)
		adminGroup.POST("/UpdateGroup", api.UpdateGroup)
		adminGroup.POST("/DeleteGroup", api.DeleteGroup)
    }
}
