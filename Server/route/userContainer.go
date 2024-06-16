package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetUserContainerRoute(r *gin.Engine) {
	userContainer := r.Group("/userContainer")
	userContainer.Use(middleware.JWTAuthMiddleware())
	{
        userContainer.POST("GetContainerListByUser", api.GetContainerListByUser)
		userContainer.POST("CreateContainerByUser", api.CreateContainerByUser)
		userContainer.POST("DestroyContainerByUser", api.DestroyContainerByUser)
		userContainer.POST("DelayContainerByUser", api.DelayContainerByUser)
	}
}
