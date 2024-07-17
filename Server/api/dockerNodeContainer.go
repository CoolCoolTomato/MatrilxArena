package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/docker"
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/localizer"
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
		response.Fail(err, localizer.GetMessage("DockerNodeContainer.InvalidArgument", c), c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = containerRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, localizer.GetMessage("DockerNodeContainer.GetDockerNodeFail", c), c)
		return
	}

	res, err := docker.GetContainerList(dockerNode)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, localizer.GetMessage("DockerNodeContainer.GetContainerListFail", c), c)
		return
	}

	response.OK(res["data"], localizer.GetMessage("DockerNodeContainer.GetContainerListSuccess", c), c)
}

func GetContainerFromDockerNode(c *gin.Context) {
	var containerRequest ContainerRequest
	err := c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 || containerRequest.DockerNodeContainerID == "" {
		response.Fail(err, localizer.GetMessage("DockerNodeContainer.InvalidArgument", c), c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = containerRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, localizer.GetMessage("DockerNodeContainer.GetDockerNodeFail", c), c)
		return
	}

	res, err := docker.GetContainer(dockerNode, containerRequest.DockerNodeContainerID)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, localizer.GetMessage("DockerNodeContainer.GetContainerFail", c), c)
		return
	}

	response.OK(res["data"], localizer.GetMessage("DockerNodeContainer.GetContainerSuccess", c), c)
}

func CreateContainerFromDockerNode(c *gin.Context) {
	var containerRequest ContainerRequest
	err := c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 || containerRequest.DockerNodeImageID == "" {
		response.Fail(err, localizer.GetMessage("DockerNodeContainer.InvalidArgument", c), c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = containerRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, localizer.GetMessage("DockerNodeContainer.GetDockerNodeFail", c), c)
		return
	}

	res, err := docker.CreateContainer(dockerNode, containerRequest.DockerNodeImageID, containerRequest.SpecifiedPorts)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, localizer.GetMessage("DockerNodeContainer.CreateContainerFail", c), c)
		return
	}

	response.OK(res["data"], localizer.GetMessage("DockerNodeContainer.CreateContainerSuccess", c), c)
}

func StartContainerFromDockerNode(c *gin.Context) {
	var containerRequest ContainerRequest
	err := c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 || containerRequest.DockerNodeContainerID == "" {
		response.Fail(err, localizer.GetMessage("DockerNodeContainer.InvalidArgument", c), c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = containerRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, localizer.GetMessage("DockerNodeContainer.GetDockerNodeFail", c), c)
		return
	}

	res, err := docker.StartContainer(dockerNode, containerRequest.DockerNodeContainerID)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, localizer.GetMessage("DockerNodeContainer.StartContainerFail", c), c)
		return
	}

	response.OK(res["data"], localizer.GetMessage("DockerNodeContainer.StartContainerSuccess", c), c)
}

func StopContainerFromDockerNode(c *gin.Context) {
	var containerRequest ContainerRequest
	err := c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 || containerRequest.DockerNodeContainerID == "" {
		response.Fail(err, localizer.GetMessage("DockerNodeContainer.InvalidArgument", c), c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = containerRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, localizer.GetMessage("DockerNodeContainer.GetDockerNodeFail", c), c)
		return
	}

	res, err := docker.StopContainer(dockerNode, containerRequest.DockerNodeContainerID)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, localizer.GetMessage("DockerNodeContainer.StopContainerFail", c), c)
		return
	}

	response.OK(res["data"], localizer.GetMessage("DockerNodeContainer.StopContainerSuccess", c), c)
}

func RemoveContainerFromDockerNode(c *gin.Context) {
	var containerRequest ContainerRequest
	err := c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 || containerRequest.DockerNodeContainerID == "" {
		response.Fail(err, localizer.GetMessage("DockerNodeContainer.InvalidArgument", c), c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = containerRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, localizer.GetMessage("DockerNodeContainer.GetDockerNodeFail", c), c)
		return
	}

	res, err := docker.RemoveContainer(dockerNode, containerRequest.DockerNodeContainerID)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, localizer.GetMessage("DockerNodeContainer.RemoveContainerFail", c), c)
		return
	}

	response.OK(res["data"], localizer.GetMessage("DockerNodeContainer.RemoveContainerSuccess", c), c)
}

func ExecuteCommandFromDockerNode(c *gin.Context) {
	var containerRequest ContainerRequest
	err := c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 || containerRequest.DockerNodeContainerID == "" || len(containerRequest.Command) == 0 {
		response.Fail(err, localizer.GetMessage("DockerNodeContainer.InvalidArgument", c), c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = containerRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, localizer.GetMessage("DockerNodeContainer.GetDockerNodeFail", c), c)
		return
	}

	res, err := docker.ExecuteCommand(dockerNode, containerRequest.DockerNodeContainerID, containerRequest.Command)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, localizer.GetMessage("DockerNodeContainer.ExecuteCommandFail", c), c)
		return
	}

	response.OK(res["data"], localizer.GetMessage("DockerNodeContainer.ExecuteCommandSuccess", c), c)
}
