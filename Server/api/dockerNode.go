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
		response.Fail(err, localizer.GetMessage("GetDockerNodeListFail", c), c)
		return
	}
	response.OK(dockerNodeList, localizer.GetMessage("GetDockerNodeListSuccess", c), c)
}

func GetDockerNode(c *gin.Context) {
	var dockerNode model.DockerNode
	err := c.ShouldBindJSON(&dockerNode)
	if err != nil || dockerNode.ID == 0 {
		response.Fail(err, localizer.GetMessage("InvalidArgument", c), c)
		return
	}

	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, localizer.GetMessage("GetDockerNodeFail", c), c)
		return
	}

	response.OK(dockerNode, localizer.GetMessage("GetDockerNodeSuccess", c), c)
}

func CreateDockerNode(c *gin.Context) {
	var dockerNode model.DockerNode
	err := c.ShouldBindJSON(&dockerNode)
	if err != nil || dockerNode.Host == "" || dockerNode.Port == "" {
		response.Fail(err, localizer.GetMessage("InvalidArgument", c), c)
		return
	}

	err = dockerNode.CreateDockerNode()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CreateDockerNodeFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("CreateDockerNodeSuccess", c), c)
}

func UpdateDockerNode(c *gin.Context) {
	var dockerNode model.DockerNode
	err := c.ShouldBindJSON(&dockerNode)
	if err != nil || dockerNode.ID == 0 || dockerNode.Host == "" || dockerNode.Port == "" {
		response.Fail(err, localizer.GetMessage("InvalidArgument", c), c)
		return
	}

	err = dockerNode.UpdateDockerNode()
	if err != nil {
		response.Fail(err, localizer.GetMessage("UpdateDockerNodeFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("UpdateDockerNodeSuccess", c), c)
}

func DeleteDockerNode(c *gin.Context) {
	var dockerNode model.DockerNode
	err := c.ShouldBindJSON(&dockerNode)
	if err != nil || dockerNode.ID == 0 {
		response.Fail(err, localizer.GetMessage("InvalidArgument", c), c)
		return
	}

	err = dockerNode.DeleteDockerNode()
	if err != nil {
		response.Fail(err, localizer.GetMessage("DeleteDockerNodeFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("DeleteDockerNodeSuccess", c), c)
}
