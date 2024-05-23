package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/gin-gonic/gin"
)

func SetChallengeRoute(r *gin.Engine) {
	challenge := r.Group("/challenge")
	{
		challenge.POST("/GetChallenge", api.GetChallenge)
		challenge.POST("/CreateChallenge", api.CreateChallenge)
		challenge.POST("/UpdateChallenge", api.UpdateChallenge)
		challenge.POST("/DeleteChallenge", api.DeleteChallenge)
	}
}