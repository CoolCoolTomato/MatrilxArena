package api

import (
	"encoding/base64"
	"fmt"
	"github.com/CoolCoolTomato/MatrilxArena/Server/docker"
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/flag"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/localizer"
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
		response.Fail(nil, localizer.GetMessage("UserContainer.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserContainer.UserNotFound", c), c)
		return
	}

	userContainers, err := manager.GetUserContainerList(username.(string), manager.StandardContainerManager{})
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserContainer.GetUserContainerListFail", c), c)
		return
	}

	response.OK(userContainers, localizer.GetMessage("UserContainer.GetUserContainerListSuccess", c), c)
}

func CreateContainerByUser(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("UserContainer.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserContainer.UserNotFound", c), c)
		return
	}

	userContainers, err := manager.GetUserContainerList(username.(string), manager.StandardContainerManager{})
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserContainer.GetUserContainerListFail", c), c)
		return
	}
	if len(userContainers) >= 3 {
		response.Fail(nil, localizer.GetMessage("UserContainer.YouCanOnlyCreateUpTo3Containers", c), c)
		return
	}

	var challenge model.Challenge
	err = c.ShouldBindJSON(&challenge)
	if err != nil || challenge.ID == 0 {
		response.Fail(err, localizer.GetMessage("UserContainer.InvalidArgument", c), c)
		return
	}

	err = challenge.GetChallenge()
	if err != nil {
		response.Fail(err, localizer.GetMessage("UserContainer.GetChallengeFail", c), c)
		return
	}

	image := challenge.Image
	var dockerNode model.DockerNode
	var dockerNodeImageID string

	dockerNodeList, err := model.GetDockerNodeList()
	if err != nil {
		response.Fail(err, localizer.GetMessage("UserContainer.GetDockerNodeListFail", c), c)
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

	res, err := docker.CreateContainer(dockerNode, dockerNodeImageID, challenge.SpecifiedPorts)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, localizer.GetMessage("UserContainer.CreateContainerFail", c), c)
		return
	}

	containerID := res["data"].(string)
	res, err = docker.GetContainer(dockerNode, containerID)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, localizer.GetMessage("UserContainer.GetContainerFail", c), c)
		return
	}

	var containerResponse ContainerResponse
	err = mapstructure.Decode(res, &containerResponse)
	if err != nil {
		response.Fail(err, localizer.GetMessage("UserContainer.DecodeContainerFail", c), c)
		return
	}

	var portMaps []manager.PortMap
	portBindings := containerResponse.Data.HostConfig.PortBindings
	for portProtocol, bindings := range portBindings {
		for _, binding := range bindings {
			portMaps = append(portMaps, manager.PortMap{
				PortProtocol: portProtocol,
				Link:         dockerNode.Address + ":" + binding.HostPort,
			})
		}
	}

	var userFlag string
	if challenge.Flag == "auto" {
		userFlag = flag.GenerateFlag()
	} else {
		userFlag = challenge.Flag
	}

	err = manager.AddUserContainer(username.(string), containerID, portMaps, dockerNode.ID, challenge.ID, userFlag, manager.StandardContainerManager{})
	if err != nil {
		response.Fail(err, localizer.GetMessage("UserContainer.AddUserContainerFail", c), c)
		return
	}

	res, err = docker.StartContainer(dockerNode, containerID)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, localizer.GetMessage("UserContainer.StartContainerFail", c), c)
		return
	}

	challenge.Commands = append(challenge.Commands, fmt.Sprintf("echo %s > /flag", userFlag))
	for _, command := range challenge.Commands {
		encodedCommand := base64.StdEncoding.EncodeToString([]byte(command))
		runCommand := fmt.Sprintf("echo %s | base64 -d | /bin/sh", encodedCommand)
		res, err = docker.ExecuteCommand(dockerNode, containerID, []string{"/bin/sh", "-c", runCommand})
		if err != nil || res["code"].(float64) != 0 {
			response.Fail(err, localizer.GetMessage("UserContainer.ExecuteCommandFail", c), c)
			return
		}
	}

	response.OK(nil, localizer.GetMessage("UserContainer.CreateContainerSuccess", c), c)
}

func DestroyContainerByUser(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("UserContainer.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserContainer.UserNotFound", c), c)
		return
	}

	var containerRequest ContainerRequest
	err = c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 || containerRequest.DockerNodeContainerID == "" {
		response.Fail(err, localizer.GetMessage("UserContainer.InvalidArgument", c), c)
		return
	}

	userContainers, err := manager.GetUserContainerList(username.(string), manager.StandardContainerManager{})
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserContainer.GetUserContainerListFail", c), c)
		return
	}

	for _, userContainer := range userContainers {
		if userContainer.DockerNodeID == containerRequest.DockerNodeID && userContainer.DockerNodeContainerID == containerRequest.DockerNodeContainerID {
			var dockerNode model.DockerNode
			dockerNode.ID = userContainer.DockerNodeID
			err = dockerNode.GetDockerNode()
			if err != nil {
				response.Fail(err, localizer.GetMessage("UserContainer.GetDockerNodeFail", c), c)
				return
			}
			err = manager.DeleteUserContainer(username.(string), userContainer.DockerNodeContainerID, manager.StandardContainerManager{})
			if err != nil {
				response.Fail(err, localizer.GetMessage("UserContainer.DeleteUserContainerFail", c), c)
				return
			}
			res, err := docker.RemoveContainer(dockerNode, userContainer.DockerNodeContainerID)
			if err != nil || res["code"].(float64) != 0 {
				response.Fail(err, localizer.GetMessage("UserContainer.RemoveContainerFail", c), c)
				return
			}
			response.OK(nil, localizer.GetMessage("UserContainer.DestroyContainerSuccess", c), c)
			return
		}
	}
	response.Fail(nil, localizer.GetMessage("UserContainer.ContainerNotFound", c), c)
}

func DelayContainerByUser(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("UserContainer.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserContainer.UserNotFound", c), c)
		return
	}

	var containerRequest ContainerRequest
	err = c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 || containerRequest.DockerNodeContainerID == "" {
		response.Fail(err, localizer.GetMessage("UserContainer.InvalidArgument", c), c)
		return
	}

	userContainers, err := manager.GetUserContainerList(username.(string), manager.StandardContainerManager{})
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserContainer.GetUserContainerListFail", c), c)
		return
	}

	for _, userContainer := range userContainers {
		if userContainer.DockerNodeID == containerRequest.DockerNodeID && userContainer.DockerNodeContainerID == containerRequest.DockerNodeContainerID {
			var dockerNode model.DockerNode
			dockerNode.ID = userContainer.DockerNodeID
			err = dockerNode.GetDockerNode()
			if err != nil {
				response.Fail(err, localizer.GetMessage("UserContainer.GetDockerNodeFail", c), c)
				return
			}
			err = manager.ResetUserContainerTime(userContainer.DockerNodeContainerID, manager.StandardContainerManager{})
			if err != nil {
				if strings.Contains(err.Error(), "remaining time is greater than 10 minutes") {
					response.Fail(err, localizer.GetMessage("UserContainer.RemainingTimeIsGreaterThan10Minutes", c), c)
					return
				} else {
					response.Fail(err, localizer.GetMessage("UserContainer.DelayContainerFail", c), c)
					return
				}
			}
			response.OK(nil, localizer.GetMessage("UserContainer.DelayContainerSuccess", c), c)
			return
		}
	}
	response.Fail(nil, localizer.GetMessage("UserContainer.ContainerNotFound", c), c)
}
