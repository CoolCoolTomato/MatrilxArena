package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/docker"
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

type ContainerRequest struct {
	DockerNodeID          uint
	DockerNodeImageID     string
	DockerNodeContainerID string
	SpecifiedPorts        []string
	Command               []string
}

func GetContainerListFromDockerNode(c *gin.Context) {
	var containerRequest ContainerRequest
	err := c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 {
		response.Fail(err, "Invalid argument", c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = containerRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, "Get dockerNode fail", c)
		return
	}

	res, err := docker.GetContainerList(dockerNode)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, "Get container list fail", c)
		return
	}

	response.OK(res["data"], "Get container list success", c)
}

func GetContainerFromDockerNode(c *gin.Context) {
	var containerRequest ContainerRequest
	err := c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 || containerRequest.DockerNodeContainerID == "" {
		response.Fail(err, "Invalid argument", c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = containerRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, "Get dockerNode fail", c)
		return
	}

	res, err := docker.GetContainer(dockerNode, containerRequest.DockerNodeContainerID)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, "Get container fail", c)
		return
	}

	response.OK(res["data"], "Get container success", c)
}

func CreateContainerFromDockerNode(c *gin.Context) {
	var containerRequest ContainerRequest
	err := c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 || containerRequest.DockerNodeImageID == "" {
		response.Fail(err, "Invalid argument", c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = containerRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, "Get dockerNode fail", c)
		return
	}

	res, err := docker.CreateContainer(dockerNode, containerRequest.DockerNodeImageID, containerRequest.SpecifiedPorts)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, "Create container fail", c)
		return
	}

	response.OK(res["data"], "Create container success", c)
}

func StartContainerFromDockerNode(c *gin.Context) {
	var containerRequest ContainerRequest
	err := c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 || containerRequest.DockerNodeContainerID == "" {
		response.Fail(err, "Invalid argument", c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = containerRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, "Get dockerNode fail", c)
		return
	}

	res, err := docker.StartContainer(dockerNode, containerRequest.DockerNodeContainerID)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, "Start container fail", c)
		return
	}

	response.OK(res["data"], "Start container success", c)
}

func StopContainerFromDockerNode(c *gin.Context) {
	var containerRequest ContainerRequest
	err := c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 || containerRequest.DockerNodeContainerID == "" {
		response.Fail(err, "Invalid argument", c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = containerRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, "Get dockerNode fail", c)
		return
	}

	res, err := docker.StopContainer(dockerNode, containerRequest.DockerNodeContainerID)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, "Stop container fail", c)
		return
	}

	response.OK(res["data"], "Stop container success", c)
}

func RemoveContainerFromDockerNode(c *gin.Context) {
	var containerRequest ContainerRequest
	err := c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 || containerRequest.DockerNodeContainerID == "" {
		response.Fail(err, "Invalid argument", c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = containerRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, "Get dockerNode fail", c)
		return
	}

	res, err := docker.RemoveContainer(dockerNode, containerRequest.DockerNodeContainerID)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, "Remove container fail", c)
		return
	}

	response.OK(res["data"], "Remove container success", c)
}

func ExecuteCommandFromDockerNode(c *gin.Context) {
	var containerRequest ContainerRequest
	err := c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 || containerRequest.DockerNodeContainerID == "" || len(containerRequest.Command) == 0 {
		response.Fail(err, "Invalid argument", c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = containerRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, "Get dockerNode fail", c)
		return
	}

	res, err := docker.ExecuteCommand(dockerNode, containerRequest.DockerNodeContainerID, containerRequest.Command)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, "Execute command fail", c)
		return
	}

	response.OK(res["data"], "Remove container success", c)
}
