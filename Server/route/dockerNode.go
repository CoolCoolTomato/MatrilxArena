package route

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/api"
	"github.com/gin-gonic/gin"
)

func SetDockerNodeRoute(r *gin.Engine) {
	dockerNode := r.Group("/dockerNode")
	{
		dockerNode.POST("/GetDockerNodeList", api.GetDockerNodeList)
		dockerNode.POST("/GetDockerNode", api.GetDockerNode)
		dockerNode.POST("/CreateDockerNode", api.CreateDockerNode)
		dockerNode.POST("/UpdateDockerNode", api.UpdateDockerNode)
		dockerNode.POST("/DeleteDockerNode", api.DeleteDockerNode)
		
		dockerNode.POST("/GetImageListFromDockerNode", api.GetImageListFromDockerNode)
		dockerNode.POST("/GetImageFromDockerNode", api.GetImageFromDockerNode)
		dockerNode.POST("/PullImageForDockerNode", api.PullImageFromDockerNode)
		dockerNode.POST("/RemoveImageFromDockerNode", api.RemoveImageFromDockerNode)

		dockerNode.POST("/GetContainerListFromDockerNode", api.GetContainerListFromDockerNode)
		dockerNode.POST("/GetContainerFromDockerNode", api.GetContainerFromDockerNode)
		dockerNode.POST("/CreateContainerFromDockerNode", api.CreateContainerFromDockerNode)
		dockerNode.POST("/StartContainerFromDockerNode", api.StartContainerFromDockerNode)
		dockerNode.POST("/StopContainerFromDockerNode", api.StopContainerFromDockerNode)
		dockerNode.POST("/RemoveContainerFromDockerNode", api.RemoveContainerFromDockerNode)
	}
}
