package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetUserGroupChallengeRoute(r *gin.Engine) {
	userChallenge := r.Group("/userGroupChallenge")
	userChallenge.Use(middleware.JWTAuthMiddleware())
	{
		userChallenge.POST("/GetUserGroupChallengeList", api.GetUserGroupChallengeList)
		userChallenge.POST("/ResetUserGroupChallenge", api.ResetUserGroupChallenge)
		userChallenge.POST("/CheckGroupChallengeFlag", api.CheckGroupChallengeFlag)
	}
}
