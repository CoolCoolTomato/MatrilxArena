package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetCTFChallengeRoute(r *gin.Engine) {
	adminCTFChallenge := r.Group("/ctfChallenge")
	adminCTFChallenge.Use(middleware.AdminAuthMiddleware())
	{
		adminCTFChallenge.POST("/GetCTFChallengeList", api.GetCTFChallengeList)
		adminCTFChallenge.POST("/GetCTFChallenge", api.GetCTFChallenge)
		adminCTFChallenge.POST("/CreateCTFChallenge", api.CreateCTFChallenge)
		adminCTFChallenge.POST("/UpdateCTFChallenge", api.UpdateCTFChallenge)
		adminCTFChallenge.POST("/DeleteCTFChallenge", api.DeleteCTFChallenge)
	}
}
