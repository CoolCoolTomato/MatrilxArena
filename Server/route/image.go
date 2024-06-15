package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/gin-gonic/gin"
)

func SetImageRoute(r *gin.Engine) {
	image := r.Group("/image")
	image.Use(middleware.AdminAuthMiddleware())
	{
		image.POST("/GetImageList", api.GetImageList)
		image.POST("/GetImage", api.GetImage)
		image.POST("/CreateImage", api.CreateImage)
		image.POST("/UpdateImage", api.UpdateImage)
		image.POST("/DeleteImage", api.DeleteImage)
	}
}
