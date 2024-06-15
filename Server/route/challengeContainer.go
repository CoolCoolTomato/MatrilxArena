package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetChallengeContainerRoute(r *gin.Engine) {
	challengeContainer := r.Group("/challengeContainer")
	challengeContainer.Use(middleware.JWTAuthMiddleware())
	{
		challengeContainer.POST("CreateContainerFromChallenge", api.CreateContainerFromChallenge)
		challengeContainer.POST("DestroyContainerByUser", api.DestroyContainerByUser)
		challengeContainer.POST("DelayContainerByUser", api.DelayContainerByUser)
	}
}
