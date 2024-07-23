package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetGroupUserRoute(r *gin.Engine) {
	adminGroupUser := r.Group("/groupUser")
	adminGroupUser.Use(middleware.AdminAuthMiddleware())
	{
		adminGroupUser.POST("/GetGroupUserList", api.GetGroupUserList)
		adminGroupUser.POST("/AddGroupUser", api.AddGroupUser)
		adminGroupUser.POST("/RemoveGroupUser", api.RemoveGroupUser)
	}
}
