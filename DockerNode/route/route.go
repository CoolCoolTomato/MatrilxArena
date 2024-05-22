package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/DockerNode/api"
	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine) {
	image := r.Group("/image")
	{
		image.POST("/GetImageList", api.GetImageList)
		image.POST("/GetImage", api.GetImage)
		image.POST("/PullImage", api.PullImage)
		image.POST("/PullImageWithAuth", api.PullImageWithAuth)
		image.POST("/RemoveImage", api.RemoveImage)
	}
	container := r.Group("/container")
	{
		container.POST("/GetContainerList", api.GetContainerList)
		container.POST("/GetContainer", api.GetContainer)
		container.POST("/CreateContainer", api.CreateContainer)
		container.POST("/StartContainer", api.StartContainer)
		container.POST("/StopContainer", api.StopContainer)
		container.POST("/RemoveContainer", api.RemoveContainer)
	}
}
