package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetUserCTFChallengeRoute(r *gin.Engine) {
	userChallenge := r.Group("/userCTFChallenge")
	userChallenge.Use(middleware.JWTAuthMiddleware())
	{
		userChallenge.POST("/GetUserCTFChallengeList", api.GetUserCTFChallengeList)
		userChallenge.POST("/ResetUserCTFChallenge", api.ResetUserCTFChallenge)
		userChallenge.POST("/CheckCTFChallengeFlag", api.CheckCTFChallengeFlag)
	}
}
