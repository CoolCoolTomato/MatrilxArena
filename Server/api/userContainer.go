package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/docker"
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/manager"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
    "github.com/mitchellh/mapstructure"
    "strings"
)

type PortBinding struct {
	HostIp   string `mapstructure:"HostIp"`
	HostPort string `mapstructure:"HostPort"`
}

type HostConfig struct {
	PortBindings map[string][]PortBinding `mapstructure:"PortBindings"`
}

type Data struct {
	HostConfig HostConfig `mapstructure:"HostConfig"`
}

type ContainerResponse struct {
	Code int    `mapstructure:"code"`
	Data Data   `mapstructure:"data"`
	Msg  string `mapstructure:"msg"`
}

func GetContainerListByUser(c *gin.Context) {
    username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, "User not found", c)
		return
	}

    userContainers, err := manager.GetUserContainers(username.(string))
	if err != nil {
		response.Fail(nil, "Get user containers fail", c)
		return
	}

    response.OK(userContainers, "Get user container success", c)
}

func CreateContainerByUser(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, "User not found", c)
		return
	}

	userContainers, err := manager.GetUserContainers(username.(string))
	if err != nil {
		response.Fail(nil, "Get user containers fail", c)
		return
	}
	if len(userContainers) >= 3 {
		response.Fail(nil, "You can only create up to 3 containers", c)
		return
	}

	var challenge model.Challenge
	err = c.ShouldBindJSON(&challenge)
	if err != nil || challenge.ID == 0 {
		response.Fail(err, "Invalid argument", c)
		return
	}

	err = challenge.GetChallenge()
	if err != nil {
		response.Fail(err, "Get challenge fail", c)
		return
	}

	image := challenge.Image
	var dockerNode model.DockerNode
	var dockerNodeImageID string

	dockerNodeList, err := model.GetDockerNodeList()
	if err != nil {
		response.Fail(err, "Get dockerNode list fail", c)
		return
	}

	for _, dockerNode = range dockerNodeList {
		res, err := docker.GetImageList(dockerNode)
		if err != nil || res["code"].(float64) != 0 {
			continue
		}
		dockerNodeImageList := res["data"].([]interface{})
		for _, dockerNodeImageRaw := range dockerNodeImageList {
			dockerNodeImage := dockerNodeImageRaw.(map[string]interface{})
			dockerNodeImageRepoTags := dockerNodeImage["RepoTags"].([]interface{})[0]
			if dockerNodeImageRepoTags.(string) == image.RepoTags {
				dockerNodeImageID = dockerNodeImage["Id"].(string)
				break
			}
		}

	}

	res, err := docker.CreateContainer(dockerNode, dockerNodeImageID)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, "Create container fail", c)
		return
	}

    containerID := res["data"].(string)
	res, err = docker.GetContainer(dockerNode, containerID)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, "Get container fail", c)
		return
	}

    var containerResponse ContainerResponse
    err = mapstructure.Decode(res, &containerResponse)
    if err != nil {
        response.Fail(err, "Decode container fail", c)
		return
    }

    var portMaps []manager.PortMap
    portBindings := containerResponse.Data.HostConfig.PortBindings
    for portProtocol, bindings := range portBindings {
        for _, binding := range bindings {
            portMaps = append(portMaps, manager.PortMap{
                PortProtocol: portProtocol,
                Link: dockerNode.Host + ":" + binding.HostPort,
            })
        }
    }

	err = manager.AddUserContainer(username.(string), containerID,  portMaps, dockerNode.ID, challenge.ID)
    if err != nil {
        response.Fail(err, "Add user container fail", c)
		return
    }

    res, err = docker.StartContainer(dockerNode, containerID)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, "Start container fail", c)
		return
	}

	response.OK(nil, "Create container success", c)
}

func DestroyContainerByUser(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, "User not found", c)
		return
	}

	var containerRequest ContainerRequest
	err := c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 || containerRequest.DockerNodeContainerID == "" {
		response.Fail(err, "Invalid argument", c)
		return
	}

	userContainers, err := manager.GetUserContainers(username.(string))
	if err != nil {
		response.Fail(nil, "Get user containers fail", c)
		return
	}

	for _, userContainer := range userContainers {
		if userContainer.DockerNodeID == containerRequest.DockerNodeID && userContainer.DockerNodeContainerID == containerRequest.DockerNodeContainerID {
			var dockerNode model.DockerNode
			dockerNode.ID = userContainer.DockerNodeID
			err = dockerNode.GetDockerNode()
			if err != nil {
				response.Fail(err, "Get dockerNode fail", c)
				return
			}
            err = manager.DeleteUserContainer(username.(string), userContainer.DockerNodeContainerID)
			if err != nil {
				response.Fail(err, "Delete user container fail", c)
				return
			}
			res, err := docker.RemoveContainer(dockerNode, userContainer.DockerNodeContainerID)
			if err != nil || res["code"].(float64) != 0 {
				response.Fail(err, "Remove container fail", c)
				return
			}
			response.OK(nil, "Destroy container success", c)
			return
		}
	}
	response.Fail(nil, "Container not found", c)
}

func DelayContainerByUser(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, "User not found", c)
		return
	}

	var containerRequest ContainerRequest
	err := c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 || containerRequest.DockerNodeContainerID == "" {
		response.Fail(err, "Invalid argument", c)
		return
	}

	userContainers, err := manager.GetUserContainers(username.(string))
	if err != nil {
		response.Fail(nil, "Get user containers fail", c)
		return
	}

	for _, userContainer := range userContainers {
		if userContainer.DockerNodeID == containerRequest.DockerNodeID && userContainer.DockerNodeContainerID == containerRequest.DockerNodeContainerID {
			var dockerNode model.DockerNode
			dockerNode.ID = userContainer.DockerNodeID
			err = dockerNode.GetDockerNode()
			if err != nil {
				response.Fail(err, "Get dockerNode fail", c)
				return
			}
			err = manager.ResetUserContainerTime(userContainer.DockerNodeContainerID)
			if err != nil {
                if strings.Contains(err.Error(), "you can only extend the container life when it has less than 10 minutes remaining") {
                    response.Fail(err, "you can only extend the container life when it has less than 10 minutes remaining", c)
				    return
                } else {
                    response.Fail(err, "Delay user container fail", c)
				    return
                }
			}
			response.OK(nil, "Delay container success", c)
			return
		}
	}
	response.Fail(nil, "Container not found", c)
}
