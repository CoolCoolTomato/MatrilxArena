package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetUserChallengeRoute(r *gin.Engine) {
	userChallenge := r.Group("/userChallenge")
	userChallenge.Use(middleware.JWTAuthMiddleware())
	{
		userChallenge.POST("GetUserChallengeList", api.GetUserChallengeList)
		userChallenge.POST("ResetUserChallenge", api.ResetUserChallenge)
		userChallenge.POST("CheckFlag", api.CheckFlag)
	}
}
