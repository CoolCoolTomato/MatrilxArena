package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetChallengeRoute(r *gin.Engine) {
	challenge := r.Group("/challenge")
	challenge.Use(middleware.JWTAuthMiddleware())
	{
		challenge.POST("/GetChallengeList", api.GetChallengeList)
		challenge.POST("/GetChallenge", api.GetChallenge)
	}
	adminChallenge := r.Group("/challenge")
	adminChallenge.Use(middleware.AdminAuthMiddleware())
	{
		adminChallenge.POST("/CreateChallenge", api.CreateChallenge)
		adminChallenge.POST("/UpdateChallenge", api.UpdateChallenge)
		adminChallenge.POST("/DeleteChallenge", api.DeleteChallenge)
	}
}
