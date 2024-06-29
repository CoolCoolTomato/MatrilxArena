package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetChallengeClassRoute(r *gin.Engine) {
	challengeClass := r.Group("/challengeClass")
	challengeClass.Use(middleware.JWTAuthMiddleware())
	{
		challengeClass.POST("/GetChallengeClassList", api.GetChallengeClassList)
		challengeClass.POST("/GetChallengeClass", api.GetChallengeClass)
	}
	adminChallengeClass := r.Group("/challengeClass")
	adminChallengeClass.Use(middleware.AdminAuthMiddleware())
	{
		adminChallengeClass.POST("/CreateChallengeClass", api.CreateChallengeClass)
		adminChallengeClass.POST("/UpdateChallengeClass", api.UpdateChallengeClass)
		adminChallengeClass.POST("/DeleteChallengeClass", api.DeleteChallengeClass)
	}
}
