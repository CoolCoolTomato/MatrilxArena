package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/DockerNode/common/response"
	"github.com/CoolCoolTomato/MatrilxArena/DockerNode/docker"
	"github.com/gin-gonic/gin"
)

type ContainerRequest struct {
	ContainerID    string
	ImageID        string
	SpecifiedPorts []string
	Command        []string
}

func GetContainerList(c *gin.Context) {
	containers, err := docker.GetContainerList()
	if err != nil {
		response.Fail(err, "Get container list fail", c)
		return
	}

	response.OK(containers, "Get container list success", c)
}

func GetContainer(c *gin.Context) {
	var containerRequest ContainerRequest
	if err := c.ShouldBindJSON(&containerRequest); err != nil {
		response.Fail(err, "Invalid argument", c)
		return
	}

	containerJSON, err := docker.GetContainer(containerRequest.ContainerID)
	if err != nil {
		response.Fail(err, "Get container fail", c)
		return
	}

	response.OK(containerJSON, "Get container success", c)
}

func CreateContainer(c *gin.Context) {
	var containerRequest ContainerRequest
	if err := c.ShouldBindJSON(&containerRequest); err != nil {
		response.Fail(err, "Invalid argument", c)
		return
	}

	containerID, err := docker.CreateContainer(containerRequest.ImageID, containerRequest.SpecifiedPorts)
	if err != nil {
		response.Fail(err, "Create container fail", c)
		return
	}

	response.OK(containerID, "Create container success", c)
}

func StartContainer(c *gin.Context) {
	var containerRequest ContainerRequest
	if err := c.ShouldBindJSON(&containerRequest); err != nil {
		response.Fail(err, "Invalid argument", c)
		return
	}

	err := docker.StartContainer(containerRequest.ContainerID)
	if err != nil {
		response.Fail(err, "Start container fail", c)
		return
	}

	response.OK(nil, "Start container success", c)
}

func StopContainer(c *gin.Context) {
	var containerRequest ContainerRequest
	if err := c.ShouldBindJSON(&containerRequest); err != nil {
		response.Fail(err, "Invalid argument", c)
		return
	}

	err := docker.StopContainer(containerRequest.ContainerID)
	if err != nil {
		response.Fail(err, "Stop container fail", c)
		return
	}

	response.OK(nil, "Stop container success", c)
}

func RemoveContainer(c *gin.Context) {
	var containerRequest ContainerRequest
	if err := c.ShouldBindJSON(&containerRequest); err != nil {
		response.Fail(err, "Invalid argument", c)
		return
	}

	err := docker.RemoveContainer(containerRequest.ContainerID)
	if err != nil {
		response.Fail(err, "Remove container fail", c)
		return
	}

	response.OK(nil, "Remove container success", c)
}

func ExecuteCommand(c *gin.Context) {
	var containerRequest ContainerRequest
	if err := c.ShouldBindJSON(&containerRequest); err != nil {
		response.Fail(err, "Invalid argument", c)
		return
	}

	err := docker.ExecuteCommand(containerRequest.ContainerID, containerRequest.Command)
	if err != nil {
		response.Fail(err, "Execute command fail", c)
		return
	}

	response.OK(nil, "Execute command success", c)
}
