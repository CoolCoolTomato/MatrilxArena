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
	if err != nil || userChallengeRequest.Flag == ""{
		response.Fail(err, "Invalid argument", c)
		return
	}

    var challenge model.Challenge
    challenge.ID = userChallengeRequest.ChallengeID
    err = challenge.GetChallenge()
    if err != nil {
        response.Fail(err, "Get challenge fail", c)
		return
    }

    if challenge.ImageID == 0 {
        ok, err := checkFlagFromChallenge(user, userChallengeRequest, challenge)
        if err != nil {
            response.Fail(err, "Check flag from challenge fail", c)
		    return
        }
        if !ok {
            response.Fail(err, "Incorrect flag", c)
		    return
        }
    } else {
        ok, err := checkFlagFromContainer(user, userChallengeRequest, challenge)
        if err != nil {
            response.Fail(err, "Check flag from container fail", c)
		    return
        }
        if !ok {
            response.Fail(err, "Incorrect flag", c)
		    return
        }
    }

	response.OK(nil, "Correct flag", c)
}

func checkFlagFromContainer(user model.User, userChallengeRequest UserChallengeRequest, challenge model.Challenge) (bool, error) {
    userContainers, err := manager.GetUserContainers(user.Username)
	if err != nil {
		return false, err
	}

	for _, userContainer := range userContainers {
		if userContainer.DockerNodeID == userChallengeRequest.DockerNodeID && userContainer.DockerNodeContainerID == userChallengeRequest.DockerNodeContainerID {
			if userChallengeRequest.Flag == userContainer.Flag {
				err = user.AddChallenge(&challenge)
				if err != nil {
					return false, err
				}
				return true, nil
			}
		}
	}
    return false, nil
}

func checkFlagFromChallenge(user model.User, userChallengeRequest UserChallengeRequest, challenge model.Challenge) (bool, error) {
    if userChallengeRequest.Flag == challenge.Flag {
        err := user.AddChallenge(&challenge)
        if err != nil {
            return false, err
        }
        return true, nil
    }
    return false, nil
}
