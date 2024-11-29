package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/challenge"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/containerManager"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/localizer"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

type UserCTFChallengeRequest struct {
	CTFChallengeID        uint
	DockerNodeID          uint
	DockerNodeContainerID string
	Flag                  string
}

func GetUserCTFChallengeList(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("UserCTFChallenge.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTFChallenge.UserNotFound", c), c)
		return
	}

	var ctf model.CTF
	err = c.ShouldBindJSON(&ctf)
	if err != nil {
		response.Fail(err, localizer.GetMessage("UserCTFChallenge.InvalidArgument", c), c)
		return
	}

	ctfChallengeList, err := user.GetCTFChallengeList(ctf)
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTFChallenge.GetUserCTFChallengeListFail", c), c)
		return
	}

	ctfChallengeList = challenge.CalculateCTFChallengeListScore(ctfChallengeList)

	response.OK(ctfChallengeList, localizer.GetMessage("UserCTFChallenge.GetUserCTFChallengeListSuccess", c), c)
}

func ResetUserCTFChallenge(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("UserCTFChallenge.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTFChallenge.UserNotFound", c), c)
		return
	}

	var userCTFChallengeRequest UserCTFChallengeRequest
	err = c.ShouldBindJSON(&userCTFChallengeRequest)
	if err != nil || userCTFChallengeRequest.CTFChallengeID == 0 {
		response.Fail(err, localizer.GetMessage("UserCTFChallenge.InvalidArgument", c), c)
		return
	}

	var ctfChallenge model.CTFChallenge
	ctfChallenge.ID = userCTFChallengeRequest.CTFChallengeID
	err = ctfChallenge.GetCTFChallenge()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTFChallenge.GetUserCTFChallengeFail", c), c)
		return
	}

	err = user.DeleteCTFChallenge(&ctfChallenge)
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTFChallenge.DeleteUserCTFChallengeFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("UserCTFChallenge.ResetCTFChallengeSuccess", c), c)
}

func CheckCTFChallengeFlag(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("UserCTFChallenge.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTFChallenge.UserNotFound", c), c)
		return
	}

	var userCTFChallengeRequest UserCTFChallengeRequest
	err = c.ShouldBindJSON(&userCTFChallengeRequest)
	if err != nil || userCTFChallengeRequest.Flag == "" {
		response.Fail(err, localizer.GetMessage("UserCTFChallenge.InvalidArgument", c), c)
		return
	}

	var ctfChallenge model.CTFChallenge
	ctfChallenge.ID = userCTFChallengeRequest.CTFChallengeID
	err = ctfChallenge.GetCTFChallenge()
	if err != nil {
		response.Fail(err, localizer.GetMessage("UserCTFChallenge.GetUserCTFChallengeFail", c), c)
		return
	}

	if ctfChallenge.ImageID == 0 {
		ok, err := checkFlagFromCTFChallenge(user, userCTFChallengeRequest, ctfChallenge)
		if err != nil {
			response.Fail(err, localizer.GetMessage("UserCTFChallenge.CheckFlagFromChallengeFail", c), c)
			return
		}
		if !ok {
			response.Fail(err, localizer.GetMessage("UserCTFChallenge.IncorrectFlag", c), c)
			return
		}
	} else {
		ok, err := checkFlagFromCTFContainer(user, userCTFChallengeRequest, ctfChallenge)
		if err != nil {
			response.Fail(err, localizer.GetMessage("UserCTFChallenge.CheckFlagFromContainerFail", c), c)
			return
		}
		if !ok {
			response.Fail(err, localizer.GetMessage("UserCTFChallenge.IncorrectFlag", c), c)
			return
		}
	}

	response.OK(nil, localizer.GetMessage("UserCTFChallenge.CorrectFlag", c), c)
}

func checkFlagFromCTFContainer(user model.User, userCTFChallengeRequest UserCTFChallengeRequest, ctfChallenge model.CTFChallenge) (bool, error) {
	userContainers, err := containerManager.GetUserContainerList(user.Username, containerManager.CTFContainerManager{})
	if err != nil {
		return false, err
	}

	for _, userContainer := range userContainers {
		if userContainer.DockerNodeID == userCTFChallengeRequest.DockerNodeID && userContainer.DockerNodeContainerID == userCTFChallengeRequest.DockerNodeContainerID {
			if userCTFChallengeRequest.Flag == userContainer.Flag {
				err = user.AddCTFChallenge(&ctfChallenge)
				if err != nil {
					return false, err
				}
				return true, nil
			}
		}
	}
	return false, nil
}

func checkFlagFromCTFChallenge(user model.User, userCTFChallengeRequest UserCTFChallengeRequest, ctfChallenge model.CTFChallenge) (bool, error) {
	if userCTFChallengeRequest.Flag == ctfChallenge.Flag {
		err := user.AddCTFChallenge(&ctfChallenge)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}
