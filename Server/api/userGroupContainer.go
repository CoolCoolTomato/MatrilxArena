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

func GetGroupContainerListByUser(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("UserGroupContainer.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroupContainer.UserNotFound", c), c)
		return
	}

	userGroupContainers, err := manager.GetUserGroupContainerList(username.(string))
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroupContainer.GetUserContainerListFail", c), c)
		return
	}

	response.OK(userGroupContainers, localizer.GetMessage("UserGroupContainer.GetUserContainerListSuccess", c), c)
}

func CreateGroupContainerByUser(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("UserGroupContainer.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroupContainer.UserNotFound", c), c)
		return
	}

	userGroupContainers, err := manager.GetUserGroupContainerList(username.(string))
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroupContainer.GetUserContainerListFail", c), c)
		return
	}
	if len(userGroupContainers) >= 3 {
		response.Fail(nil, localizer.GetMessage("UserGroupContainer.YouCanOnlyCreateUpTo3Containers", c), c)
		return
	}

	var groupChallenge model.GroupChallenge
	err = c.ShouldBindJSON(&groupChallenge)
	if err != nil || groupChallenge.ID == 0 {
		response.Fail(err, localizer.GetMessage("UserGroupContainer.InvalidArgument", c), c)
		return
	}

	err = groupChallenge.GetGroupChallenge()
	if err != nil {
		response.Fail(err, localizer.GetMessage("UserGroupContainer.GetGroupChallengeFail", c), c)
		return
	}

	image := groupChallenge.Image
	var dockerNode model.DockerNode
	var dockerNodeImageID string

	dockerNodeList, err := model.GetDockerNodeList()
	if err != nil {
		response.Fail(err, localizer.GetMessage("UserGroupContainer.GetDockerNodeListFail", c), c)
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

	res, err := docker.CreateContainer(dockerNode, dockerNodeImageID, groupChallenge.SpecifiedPorts)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, localizer.GetMessage("UserGroupContainer.CreateContainerFail", c), c)
		return
	}

	containerID := res["data"].(string)
	res, err = docker.GetContainer(dockerNode, containerID)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, localizer.GetMessage("UserGroupContainer.GetContainerFail", c), c)
		return
	}

	var containerResponse ContainerResponse
	err = mapstructure.Decode(res, &containerResponse)
	if err != nil {
		response.Fail(err, localizer.GetMessage("UserGroupContainer.DecodeContainerFail", c), c)
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
	if groupChallenge.Flag == "auto" {
		userFlag = flag.GenerateFlag()
	} else {
		userFlag = groupChallenge.Flag
	}

	err = manager.AddUserGroupContainer(username.(string), containerID, portMaps, dockerNode.ID, groupChallenge.ID, userFlag)
	if err != nil {
		response.Fail(err, localizer.GetMessage("UserGroupContainer.AddUserContainerFail", c), c)
		return
	}

	res, err = docker.StartContainer(dockerNode, containerID)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, localizer.GetMessage("UserGroupContainer.StartContainerFail", c), c)
		return
	}

	groupChallenge.Commands = append(groupChallenge.Commands, fmt.Sprintf("echo %s > /flag", userFlag))
	for _, command := range groupChallenge.Commands {
		encodedCommand := base64.StdEncoding.EncodeToString([]byte(command))
		runCommand := fmt.Sprintf("echo %s | base64 -d | /bin/sh", encodedCommand)
		res, err = docker.ExecuteCommand(dockerNode, containerID, []string{"/bin/sh", "-c", runCommand})
		if err != nil || res["code"].(float64) != 0 {
			response.Fail(err, localizer.GetMessage("UserGroupContainer.ExecuteCommandFail", c), c)
			return
		}
	}

	response.OK(nil, localizer.GetMessage("UserGroupContainer.CreateContainerSuccess", c), c)
}

func DestroyGroupContainerByUser(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("UserGroupContainer.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroupContainer.UserNotFound", c), c)
		return
	}

	var containerRequest ContainerRequest
	err = c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 || containerRequest.DockerNodeContainerID == "" {
		response.Fail(err, localizer.GetMessage("UserGroupContainer.InvalidArgument", c), c)
		return
	}

	userGroupContainers, err := manager.GetUserGroupContainerList(username.(string))
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroupContainer.GetUserContainerListFail", c), c)
		return
	}

	for _, userGroupContainer := range userGroupContainers {
		if userGroupContainer.DockerNodeID == containerRequest.DockerNodeID && userGroupContainer.DockerNodeContainerID == containerRequest.DockerNodeContainerID {
			var dockerNode model.DockerNode
			dockerNode.ID = userGroupContainer.DockerNodeID
			err = dockerNode.GetDockerNode()
			if err != nil {
				response.Fail(err, localizer.GetMessage("UserGroupContainer.GetDockerNodeFail", c), c)
				return
			}
			err = manager.DeleteUserGroupContainer(username.(string), userGroupContainer.DockerNodeContainerID)
			if err != nil {
				response.Fail(err, localizer.GetMessage("UserGroupContainer.DeleteUserContainerFail", c), c)
				return
			}
			res, err := docker.RemoveContainer(dockerNode, userGroupContainer.DockerNodeContainerID)
			if err != nil || res["code"].(float64) != 0 {
				response.Fail(err, localizer.GetMessage("UserGroupContainer.RemoveContainerFail", c), c)
				return
			}
			response.OK(nil, localizer.GetMessage("UserGroupContainer.DestroyContainerSuccess", c), c)
			return
		}
	}
	response.Fail(nil, localizer.GetMessage("UserGroupContainer.ContainerNotFound", c), c)
}

func DelayGroupContainerByUser(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("UserGroupContainer.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroupContainer.UserNotFound", c), c)
		return
	}

	var containerRequest ContainerRequest
	err = c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 || containerRequest.DockerNodeContainerID == "" {
		response.Fail(err, localizer.GetMessage("UserGroupContainer.InvalidArgument", c), c)
		return
	}

	userGroupContainers, err := manager.GetUserGroupContainerList(username.(string))
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroupContainer.GetUserContainerListFail", c), c)
		return
	}

	for _, userGroupContainer := range userGroupContainers {
		if userGroupContainer.DockerNodeID == containerRequest.DockerNodeID && userGroupContainer.DockerNodeContainerID == containerRequest.DockerNodeContainerID {
			var dockerNode model.DockerNode
			dockerNode.ID = userGroupContainer.DockerNodeID
			err = dockerNode.GetDockerNode()
			if err != nil {
				response.Fail(err, localizer.GetMessage("UserGroupContainer.GetDockerNodeFail", c), c)
				return
			}
			err = manager.ResetUserGroupContainerTime(userGroupContainer.DockerNodeContainerID)
			if err != nil {
				if strings.Contains(err.Error(), "remaining time is greater than 10 minutes") {
					response.Fail(err, localizer.GetMessage("UserGroupContainer.RemainingTimeIsGreaterThan10Minutes", c), c)
					return
				} else {
					response.Fail(err, localizer.GetMessage("UserGroupContainer.DelayContainerFail", c), c)
					return
				}
			}
			response.OK(nil, localizer.GetMessage("UserGroupContainer.DelayContainerSuccess", c), c)
			return
		}
	}
	response.Fail(nil, localizer.GetMessage("UserGroupContainer.ContainerNotFound", c), c)
}
