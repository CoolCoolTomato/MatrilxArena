package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetGroupChallengeRoute(r *gin.Engine) {
	groupChallenge := r.Group("/groupChallenge")
	groupChallenge.Use(middleware.JWTAuthMiddleware())
	{
		groupChallenge.POST("/GetGroupChallengeList", api.GetGroupChallengeList)
        groupChallenge.POST("/GetGroupChallengeListByQuery", api.GetGroupChallengeListByQuery)
		groupChallenge.POST("/GetGroupChallenge", api.GetGroupChallenge)
	}
	adminGroupChallenge := r.Group("/groupChallenge")
	adminGroupChallenge.Use(middleware.AdminAuthMiddleware())
	{
		adminGroupChallenge.POST("/CreateGroupChallenge", api.CreateGroupChallenge)
		adminGroupChallenge.POST("/UpdateGroupChallenge", api.UpdateGroupChallenge)
		adminGroupChallenge.POST("/DeleteGroupChallenge", api.DeleteGroupChallenge)
	}
}
