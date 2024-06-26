package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/manager"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

type CheckFlagRequest struct {
	DockerNodeID          uint
	DockerNodeContainerID string
	Flag                  string
}

func GetUserChallenge(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, "Invalid token", c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, "User not found", c)
		return
	}

	challengeList, err := user.GetChallengeList()
	if err != nil {
		response.Fail(nil, "Get challenge list fail", c)
		return
	}

	response.OK(challengeList, "Get challenge list success", c)
}

func CheckFlag(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, "Invalid token", c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, "User not found", c)
		return
	}

	var checkFlagRequest CheckFlagRequest
	err = c.ShouldBindJSON(&checkFlagRequest)
	if err != nil || checkFlagRequest.DockerNodeID == 0 || checkFlagRequest.DockerNodeContainerID == "" {
		response.Fail(err, "Invalid argument", c)
		return
	}

	userContainers, err := manager.GetUserContainers(username.(string))
	if err != nil {
		response.Fail(nil, "Get user containers fail", c)
		return
	}

	for _, userContainer := range userContainers {
		if userContainer.DockerNodeID == checkFlagRequest.DockerNodeID && userContainer.DockerNodeContainerID == checkFlagRequest.DockerNodeContainerID {
			if checkFlagRequest.Flag == userContainer.Flag {
				response.OK(nil, "Correct flag", c)
				return
			}
		}
	}
	response.Fail(nil, "Incorrect flag", c)
}
