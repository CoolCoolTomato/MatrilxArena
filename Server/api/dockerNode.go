package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/common/response"
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/gin-gonic/gin"
)

func GetDockerNodeList(c *gin.Context) {
	dockerNodeList, err := model.GetDockerNodeList()
	if err != nil {
		response.Fail(err, "Get dockerNode list fail", c)
		return
	}
	response.OK(dockerNodeList, "Get dockerNode list success", c)
}

func GetDockerNode(c *gin.Context) {
	var dockerNode model.DockerNode
	err := c.ShouldBindJSON(&dockerNode)
	if err != nil || dockerNode.ID == 0 {
		response.Fail(err, "Invalid argument", c)
		return
	}
	
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, "Get dockerNode fail", c)
		return
	}
	
	response.OK(dockerNode, "Get dockerNode success", c)
}

func CreateDockerNode(c *gin.Context) {
	var dockerNode model.DockerNode
	err := c.ShouldBindJSON(&dockerNode)
	if err != nil {
		response.Fail(err, "Invalid argument", c)
		return
	}
	
	err = dockerNode.CreateDockerNode()
	if err != nil {
		response.Fail(err, "Create dockerNode fail", c)
		return
	}
	
	response.OK(nil, "Create dockerNode success", c)
}

func UpdateDockerNode(c *gin.Context) {
	var dockerNode model.DockerNode
	err := c.ShouldBindJSON(&dockerNode)
	if err != nil || dockerNode.ID == 0 {
		response.Fail(err, "Invalid argument", c)
		return
	}
	
	err = dockerNode.UpdateDockerNode()
	if err != nil {
		response.Fail(err, "Update dockerNode fail", c)
		return
	}
	
	response.OK(nil, "Update dockerNode success", c)
}

func DeleteDockerNode(c *gin.Context) {
	var dockerNode model.DockerNode
	err := c.ShouldBindJSON(&dockerNode)
	if err != nil || dockerNode.ID == 0 {
		response.Fail(err, "Invalid argument", c)
		return
	}
	
	err = dockerNode.DeleteDockerNode()
	if err != nil {
		response.Fail(err, "Delete dockerNode fail", c)
		return
	}
	
	response.OK(nil, "Delete dockerNode success", c)
}

