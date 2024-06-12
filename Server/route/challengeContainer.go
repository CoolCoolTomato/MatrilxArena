package route

import (
    "github.com/CoolCoolTomato/MatrilxArena/Server/api"
    "github.com/gin-gonic/gin"
)

func SetChallengeContainerRoute(r *gin.Engine) {
    challengeContainer := r.Group("/challengeContainer")
    {
        challengeContainer.POST("CreateContainerFromChallenge", api.CreateContainerFromChallenge)
    }
}
