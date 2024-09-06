package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetUserCTFContainerRoute(r *gin.Engine) {
	userContainer := r.Group("/userCTFContainer")
	userContainer.Use(middleware.JWTAuthMiddleware())
	{
		userContainer.POST("/GetCTFContainerListByUser", api.GetCTFContainerListByUser)
		userContainer.POST("/CreateCTFContainerByUser", api.CreateCTFContainerByUser)
		userContainer.POST("/DestroyCTFContainerByUser", api.DestroyCTFContainerByUser)
		userContainer.POST("/DelayCTFContainerByUser", api.DelayCTFContainerByUser)
	}
}
