package api

import (
    "encoding/base64"
    "fmt"
    "github.com/CoolCoolTomato/MatrilxArena/Server/docker"
    "github.com/CoolCoolTomato/MatrilxArena/Server/model"
    "github.com/CoolCoolTomato/MatrilxArena/Server/utils/containerManager"
    "github.com/CoolCoolTomato/MatrilxArena/Server/utils/flag"
    "github.com/CoolCoolTomato/MatrilxArena/Server/utils/localizer"
    "github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
    "github.com/gin-gonic/gin"
    "github.com/mitchellh/mapstructure"
    "strings"
)

func GetCTFContainerListByUser(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("UserCTFContainer.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTFContainer.UserNotFound", c), c)
		return
	}

	userCTFContainers, err := containerManager.GetUserContainerList(username.(string), containerManager.CTFContainerManager{})
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTFContainer.GetUserContainerListFail", c), c)
		return
	}

	response.OK(userCTFContainers, localizer.GetMessage("UserCTFContainer.GetUserContainerListSuccess", c), c)
}

func CreateCTFContainerByUser(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("UserCTFContainer.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTFContainer.UserNotFound", c), c)
		return
	}

	userCTFContainers, err := containerManager.GetUserContainerList(username.(string), containerManager.CTFContainerManager{})
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTFContainer.GetUserContainerListFail", c), c)
		return
	}
	if len(userCTFContainers) >= 3 {
		response.Fail(nil, localizer.GetMessage("UserCTFContainer.YouCanOnlyCreateUpTo3Containers", c), c)
		return
	}

	var ctfChallenge model.CTFChallenge
	err = c.ShouldBindJSON(&ctfChallenge)
	if err != nil || ctfChallenge.ID == 0 {
		response.Fail(err, localizer.GetMessage("UserCTFContainer.InvalidArgument", c), c)
		return
	}

	err = ctfChallenge.GetCTFChallenge()
	if err != nil {
		response.Fail(err, localizer.GetMessage("UserCTFContainer.GetCTFChallengeFail", c), c)
		return
	}

	image := ctfChallenge.Image
	var dockerNode model.DockerNode
	var dockerNodeImageID string

	dockerNodeList, err := model.GetDockerNodeList()
	if err != nil {
		response.Fail(err, localizer.GetMessage("UserCTFContainer.GetDockerNodeListFail", c), c)
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

	res, err := docker.CreateContainer(dockerNode, dockerNodeImageID, ctfChallenge.SpecifiedPorts)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, localizer.GetMessage("UserCTFContainer.CreateContainerFail", c), c)
		return
	}

	containerID := res["data"].(string)
	res, err = docker.GetContainer(dockerNode, containerID)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, localizer.GetMessage("UserCTFContainer.GetContainerFail", c), c)
		return
	}

	var containerResponse ContainerResponse
	err = mapstructure.Decode(res, &containerResponse)
	if err != nil {
		response.Fail(err, localizer.GetMessage("UserCTFContainer.DecodeContainerFail", c), c)
		return
	}

	var portMaps []containerManager.PortMap
	portBindings := containerResponse.Data.HostConfig.PortBindings
	for portProtocol, bindings := range portBindings {
		for _, binding := range bindings {
			portMaps = append(portMaps, containerManager.PortMap{
				PortProtocol: portProtocol,
				Link:         dockerNode.Address + ":" + binding.HostPort,
			})
		}
	}

	var userFlag string
	if ctfChallenge.Flag == "auto" {
		userFlag = flag.GenerateFlag()
	} else {
		userFlag = ctfChallenge.Flag
	}

	err = containerManager.AddUserContainer(username.(string), containerID, portMaps, dockerNode.ID, ctfChallenge.ID, userFlag, containerManager.CTFContainerManager{})
	if err != nil {
		response.Fail(err, localizer.GetMessage("UserCTFContainer.AddUserContainerFail", c), c)
		return
	}

	res, err = docker.StartContainer(dockerNode, containerID)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, localizer.GetMessage("UserCTFContainer.StartContainerFail", c), c)
		return
	}

	ctfChallenge.Commands = append(ctfChallenge.Commands, fmt.Sprintf("echo %s > /flag", userFlag))
	for _, command := range ctfChallenge.Commands {
		encodedCommand := base64.StdEncoding.EncodeToString([]byte(command))
		runCommand := fmt.Sprintf("echo %s | base64 -d | /bin/sh", encodedCommand)
		res, err = docker.ExecuteCommand(dockerNode, containerID, []string{"/bin/sh", "-c", runCommand})
		if err != nil || res["code"].(float64) != 0 {
			response.Fail(err, localizer.GetMessage("UserCTFContainer.ExecuteCommandFail", c), c)
			return
		}
	}

	response.OK(nil, localizer.GetMessage("UserCTFContainer.CreateContainerSuccess", c), c)
}

func DestroyCTFContainerByUser(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("UserCTFContainer.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTFContainer.UserNotFound", c), c)
		return
	}

	var containerRequest ContainerRequest
	err = c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 || containerRequest.DockerNodeContainerID == "" {
		response.Fail(err, localizer.GetMessage("UserCTFContainer.InvalidArgument", c), c)
		return
	}

	userCTFContainers, err := containerManager.GetUserContainerList(username.(string), containerManager.CTFContainerManager{})
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTFContainer.GetUserContainerListFail", c), c)
		return
	}

	for _, userCTFContainer := range userCTFContainers {
		if userCTFContainer.DockerNodeID == containerRequest.DockerNodeID && userCTFContainer.DockerNodeContainerID == containerRequest.DockerNodeContainerID {
			var dockerNode model.DockerNode
			dockerNode.ID = userCTFContainer.DockerNodeID
			err = dockerNode.GetDockerNode()
			if err != nil {
				response.Fail(err, localizer.GetMessage("UserCTFContainer.GetDockerNodeFail", c), c)
				return
			}
			err = containerManager.DeleteUserContainer(username.(string), userCTFContainer.DockerNodeContainerID, containerManager.CTFContainerManager{})
			if err != nil {
				response.Fail(err, localizer.GetMessage("UserCTFContainer.DeleteUserContainerFail", c), c)
				return
			}
			res, err := docker.RemoveContainer(dockerNode, userCTFContainer.DockerNodeContainerID)
			if err != nil || res["code"].(float64) != 0 {
				response.Fail(err, localizer.GetMessage("UserCTFContainer.RemoveContainerFail", c), c)
				return
			}
			response.OK(nil, localizer.GetMessage("UserCTFContainer.DestroyContainerSuccess", c), c)
			return
		}
	}
	response.Fail(nil, localizer.GetMessage("UserCTFContainer.ContainerNotFound", c), c)
}

func DelayCTFContainerByUser(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("UserCTFContainer.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTFContainer.UserNotFound", c), c)
		return
	}

	var containerRequest ContainerRequest
	err = c.ShouldBindJSON(&containerRequest)
	if err != nil || containerRequest.DockerNodeID == 0 || containerRequest.DockerNodeContainerID == "" {
		response.Fail(err, localizer.GetMessage("UserCTFContainer.InvalidArgument", c), c)
		return
	}

	userCTFContainers, err := containerManager.GetUserContainerList(username.(string), containerManager.CTFContainerManager{})
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTFContainer.GetUserContainerListFail", c), c)
		return
	}

	for _, userCTFContainer := range userCTFContainers {
		if userCTFContainer.DockerNodeID == containerRequest.DockerNodeID && userCTFContainer.DockerNodeContainerID == containerRequest.DockerNodeContainerID {
			var dockerNode model.DockerNode
			dockerNode.ID = userCTFContainer.DockerNodeID
			err = dockerNode.GetDockerNode()
			if err != nil {
				response.Fail(err, localizer.GetMessage("UserCTFContainer.GetDockerNodeFail", c), c)
				return
			}
			err = containerManager.ResetUserContainerTime(userCTFContainer.DockerNodeContainerID, containerManager.CTFContainerManager{})
			if err != nil {
				if strings.Contains(err.Error(), "remaining time is greater than 10 minutes") {
					response.Fail(err, localizer.GetMessage("UserCTFContainer.RemainingTimeIsGreaterThan10Minutes", c), c)
					return
				} else {
					response.Fail(err, localizer.GetMessage("UserCTFContainer.DelayContainerFail", c), c)
					return
				}
			}
			response.OK(nil, localizer.GetMessage("UserCTFContainer.DelayContainerSuccess", c), c)
			return
		}
	}
	response.Fail(nil, localizer.GetMessage("UserCTFContainer.ContainerNotFound", c), c)
}
