package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/manager"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

type UserChallengeRequest struct {
	ChallengeID           uint
	DockerNodeID          uint
	DockerNodeContainerID string
	Flag                  string
}

func GetUserChallengeList(c *gin.Context) {
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

func ResetUserChallenge(c *gin.Context) {
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

	var userChallengeRequest UserChallengeRequest
	err = c.ShouldBindJSON(&userChallengeRequest)
	if err != nil || userChallengeRequest.ChallengeID == 0 {
		response.Fail(err, "Invalid argument", c)
		return
	}

	var challenge model.Challenge
	challenge.ID = userChallengeRequest.ChallengeID
	err = challenge.GetChallenge()
	if err != nil {
		response.Fail(nil, "Get challenge fail", c)
		return
	}

	err = user.DeleteChallenge(&challenge)
	if err != nil {
		response.Fail(nil, "Delete challenge fail", c)
		return
	}

	response.OK(nil, "Reset challenge success", c)
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

	var userChallengeRequest UserChallengeRequest
	err = c.ShouldBindJSON(&userChallengeRequest)
	if err != nil || userChallengeRequest.DockerNodeID == 0 || userChallengeRequest.DockerNodeContainerID == "" {
		response.Fail(err, "Invalid argument", c)
		return
	}

	userContainers, err := manager.GetUserContainers(username.(string))
	if err != nil {
		response.Fail(nil, "Get user containers fail", c)
		return
	}

	for _, userContainer := range userContainers {
		if userContainer.DockerNodeID == userChallengeRequest.DockerNodeID && userContainer.DockerNodeContainerID == userChallengeRequest.DockerNodeContainerID {
			if userChallengeRequest.Flag == userContainer.Flag {
				challengeID := userContainer.ChallengeID
				var challenge model.Challenge
				challenge.ID = challengeID
				err = challenge.GetChallenge()
				if err != nil {
					response.Fail(nil, "Get challenge fail", c)
					return
				}
				err = user.AddChallenge(&challenge)
				if err != nil {
					response.Fail(nil, "Add challenge fail", c)
					return
				}
				response.OK(nil, "Correct flag", c)
				return
			}
		}
	}
	response.Fail(nil, "Incorrect flag", c)
}
