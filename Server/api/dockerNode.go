package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/localizer"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

func GetDockerNodeList(c *gin.Context) {
	dockerNodeList, err := model.GetDockerNodeList()
	if err != nil {
		response.Fail(err, localizer.GetMessage("DockerNode.GetDockerNodeListFail", c), c)
		return
	}
	response.OK(dockerNodeList, localizer.GetMessage("DockerNode.GetDockerNodeListSuccess", c), c)
}

func GetDockerNode(c *gin.Context) {
	var dockerNode model.DockerNode
	err := c.ShouldBindJSON(&dockerNode)
	if err != nil || dockerNode.ID == 0 {
		response.Fail(err, localizer.GetMessage("DockerNode.InvalidArgument", c), c)
		return
	}

	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, localizer.GetMessage("DockerNode.GetDockerNodeFail", c), c)
		return
	}

	response.OK(dockerNode, localizer.GetMessage("DockerNode.GetDockerNodeSuccess", c), c)
}

func CreateDockerNode(c *gin.Context) {
	var dockerNode model.DockerNode
	err := c.ShouldBindJSON(&dockerNode)
	if err != nil || dockerNode.Host == "" || dockerNode.Port == "" {
		response.Fail(err, localizer.GetMessage("DockerNode.InvalidArgument", c), c)
		return
	}

	err = dockerNode.CreateDockerNode()
	if err != nil {
		response.Fail(err, localizer.GetMessage("DockerNode.CreateDockerNodeFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("DockerNode.CreateDockerNodeSuccess", c), c)
}

func UpdateDockerNode(c *gin.Context) {
	var dockerNode model.DockerNode
	err := c.ShouldBindJSON(&dockerNode)
	if err != nil || dockerNode.ID == 0 || dockerNode.Host == "" || dockerNode.Port == "" {
		response.Fail(err, localizer.GetMessage("DockerNode.InvalidArgument", c), c)
		return
	}

	err = dockerNode.UpdateDockerNode()
	if err != nil {
		response.Fail(err, localizer.GetMessage("DockerNode.UpdateDockerNodeFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("DockerNode.UpdateDockerNodeSuccess", c), c)
}

func DeleteDockerNode(c *gin.Context) {
	var dockerNode model.DockerNode
	err := c.ShouldBindJSON(&dockerNode)
	if err != nil || dockerNode.ID == 0 {
		response.Fail(err, localizer.GetMessage("DockerNode.InvalidArgument", c), c)
		return
	}

	err = dockerNode.DeleteDockerNode()
	if err != nil {
		response.Fail(err, localizer.GetMessage("DockerNode.DeleteDockerNodeFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("DockerNode.DeleteDockerNodeSuccess", c), c)
}
