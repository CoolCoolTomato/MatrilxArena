package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetCTFTeamRoute(r *gin.Engine) {
	adminCTFTeam := r.Group("/ctfTeam")
	adminCTFTeam.Use(middleware.AdminAuthMiddleware())
	{
		adminCTFTeam.POST("/GetCTFTeamList", api.GetCTFTeamList)
		adminCTFTeam.POST("/GetCTFTeam", api.GetCTFTeam)
		adminCTFTeam.POST("/CreateCTFTeam", api.CreateCTFTeam)
		adminCTFTeam.POST("/UpdateCTFTeam", api.UpdateCTFTeam)
		adminCTFTeam.POST("/DeleteCTFTeam", api.DeleteCTFTeam)
	}
}
