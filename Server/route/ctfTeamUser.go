package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetCTFTeamUserRoute(r *gin.Engine) {
	adminCTFTeamUser := r.Group("/ctfTeamUser")
	adminCTFTeamUser.Use(middleware.AdminAuthMiddleware())
	{
		adminCTFTeamUser.POST("/GetCTFTeamUserList", api.GetCTFTeamUserList)
		adminCTFTeamUser.POST("/AddCTFTeamUser", api.AddCTFTeamUser)
		adminCTFTeamUser.POST("/RemoveCTFTeamUser", api.RemoveCTFTeamUser)
	}
}
