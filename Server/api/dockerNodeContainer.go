package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/common/response"
	"github.com/CoolCoolTomato/MatrilxArena/Server/docker"
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/gin-gonic/gin"
)

type ContainerRequest struct {
	DockerNodeID uint
	DockerNodeImageID string
	DockerNodeContainerID string
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
	
	url := "http://" + dockerNode.Host + ":" + dockerNode.Port + docker.PathGetContainerList
	res, err := docker.GetContainerList(url)
	if err != nil || res["code"].(float64) != 0{
		response.Fail(err, "Get container list fail", c)
		return
	}

	response.OK(res["data"], "Get container list success", c)
}

func GetContainerFromDockerNode(c *gin.Context) {
	var containerRequest ContainerRequest
	err := c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 || containerRequest.DockerNodeContainerID == ""{
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
	
	url := "http://" + dockerNode.Host + ":" + dockerNode.Port + docker.PathGetContainer
	res, err := docker.GetContainer(url, containerRequest.DockerNodeContainerID)
	if err != nil || res["code"].(float64) != 0{
		response.Fail(err, "Get container fail", c)
		return
	}

	response.OK(res["data"], "Get container success", c)
}

func CreateContainerFromDockerNode(c *gin.Context) {
	var containerRequest ContainerRequest
	err := c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 || containerRequest.DockerNodeImageID == ""{
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
	
	url := "http://" + dockerNode.Host + ":" + dockerNode.Port + docker.PathCreateContainer
	res, err := docker.CreateContainer(url, containerRequest.DockerNodeImageID)
	if err != nil || res["code"].(float64) != 0{
		response.Fail(err, "Create container fail", c)
		return
	}

	response.OK(res["data"], "Create container success", c)
}

func StartContainerFromDockerNode(c *gin.Context) {
	var containerRequest ContainerRequest
	err := c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 || containerRequest.DockerNodeContainerID == ""{
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
	
	url := "http://" + dockerNode.Host + ":" + dockerNode.Port + docker.PathStartContainer
	res, err := docker.StartContainer(url, containerRequest.DockerNodeContainerID)
	if err != nil || res["code"].(float64) != 0{
		response.Fail(err, "Start container fail", c)
		return
	}

	response.OK(res["data"], "Start container success", c)
}

func StopContainerFromDockerNode(c *gin.Context) {
	var containerRequest ContainerRequest
	err := c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 || containerRequest.DockerNodeContainerID == ""{
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
	
	url := "http://" + dockerNode.Host + ":" + dockerNode.Port + docker.PathStopContainer
	res, err := docker.StopContainer(url, containerRequest.DockerNodeContainerID)
	if err != nil || res["code"].(float64) != 0{
		response.Fail(err, "Stop container fail", c)
		return
	}

	response.OK(res["data"], "Stop container success", c)
}

func RemoveContainerFromDockerNode(c *gin.Context) {
	var containerRequest ContainerRequest
	err := c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 || containerRequest.DockerNodeContainerID == ""{
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
	
	url := "http://" + dockerNode.Host + ":" + dockerNode.Port + docker.PathRemoveContainer
	res, err := docker.RemoveContainer(url, containerRequest.DockerNodeContainerID)
	if err != nil || res["code"].(float64) != 0{
		response.Fail(err, "Remove container fail", c)
		return
	}

	response.OK(res["data"], "Remove container success", c)
}