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
		dockerNode.POST("/PullImage", api.PullImage)
	}
}