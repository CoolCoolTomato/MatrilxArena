package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/localizer"
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
		response.Fail(nil, localizer.GetMessage("InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserNotFound", c), c)
		return
	}

	challengeList, err := user.GetChallengeList()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("GetUserChallengeListFail", c), c)
		return
	}

	response.OK(challengeList, localizer.GetMessage("GetUserChallengeListSuccess", c), c)
}

func ResetUserChallenge(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserNotFound", c), c)
		return
	}

	var userChallengeRequest UserChallengeRequest
	err = c.ShouldBindJSON(&userChallengeRequest)
	if err != nil || userChallengeRequest.ChallengeID == 0 {
		response.Fail(err, localizer.GetMessage("InvalidArgument", c), c)
		return
	}

	var challenge model.Challenge
	challenge.ID = userChallengeRequest.ChallengeID
	err = challenge.GetChallenge()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("GetUserChallengeFail", c), c)
		return
	}

	err = user.DeleteChallenge(&challenge)
	if err != nil {
		response.Fail(nil, localizer.GetMessage("DeleteUserChallengeFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("ResetChallengeSuccess", c), c)
}

func CheckFlag(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserNotFound", c), c)
		return
	}

	var userChallengeRequest UserChallengeRequest
	err = c.ShouldBindJSON(&userChallengeRequest)
	if err != nil || userChallengeRequest.Flag == "" {
		response.Fail(err, localizer.GetMessage("InvalidArgument", c), c)
		return
	}

	var challenge model.Challenge
	challenge.ID = userChallengeRequest.ChallengeID
	err = challenge.GetChallenge()
	if err != nil {
		response.Fail(err, localizer.GetMessage("GetUserChallengeFail", c), c)
		return
	}

	if challenge.ImageID == 0 {
		ok, err := checkFlagFromChallenge(user, userChallengeRequest, challenge)
		if err != nil {
			response.Fail(err, localizer.GetMessage("CheckFlagFromChallengeFail", c), c)
			return
		}
		if !ok {
			response.Fail(err, localizer.GetMessage("IncorrectFlag", c), c)
			return
		}
	} else {
		ok, err := checkFlagFromContainer(user, userChallengeRequest, challenge)
		if err != nil {
			response.Fail(err, localizer.GetMessage("CheckFlagFromContainerFail", c), c)
			return
		}
		if !ok {
			response.Fail(err, localizer.GetMessage("IncorrectFlag", c), c)
			return
		}
	}

	response.OK(nil, localizer.GetMessage("CorrectFlag", c), c)
}

func checkFlagFromContainer(user model.User, userChallengeRequest UserChallengeRequest, challenge model.Challenge) (bool, error) {
	userContainers, err := manager.GetUserContainerList(user.Username)
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
