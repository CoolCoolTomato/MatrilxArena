package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/localizer"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/manager"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

type UserGroupChallengeRequest struct {
	GroupChallengeID      uint
	DockerNodeID          uint
	DockerNodeContainerID string
	Flag                  string
}

func GetUserGroupChallengeList(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("UserGroupChallenge.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroupChallenge.UserNotFound", c), c)
		return
	}

	groupChallengeList, err := user.GetGroupChallengeList()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroupChallenge.GetUserGroupChallengeListFail", c), c)
		return
	}

	response.OK(groupChallengeList, localizer.GetMessage("UserGroupChallenge.GetUserGroupChallengeListSuccess", c), c)
}

func ResetUserGroupChallenge(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("UserGroupChallenge.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroupChallenge.UserNotFound", c), c)
		return
	}

	var userGroupChallengeRequest UserGroupChallengeRequest
	err = c.ShouldBindJSON(&userGroupChallengeRequest)
	if err != nil || userGroupChallengeRequest.GroupChallengeID == 0 {
		response.Fail(err, localizer.GetMessage("UserGroupChallenge.InvalidArgument", c), c)
		return
	}

	var groupChallenge model.GroupChallenge
	groupChallenge.ID = userGroupChallengeRequest.GroupChallengeID
	err = groupChallenge.GetGroupChallenge()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroupChallenge.GetUserGroupChallengeFail", c), c)
		return
	}

	err = user.DeleteGroupChallenge(&groupChallenge)
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroupChallenge.DeleteUserGroupChallengeFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("UserGroupChallenge.ResetGroupChallengeSuccess", c), c)
}

func CheckGroupChallengeFlag(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("UserGroupChallenge.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroupChallenge.UserNotFound", c), c)
		return
	}

	var userGroupChallengeRequest UserGroupChallengeRequest
	err = c.ShouldBindJSON(&userGroupChallengeRequest)
	if err != nil || userGroupChallengeRequest.Flag == "" {
		response.Fail(err, localizer.GetMessage("UserGroupChallenge.InvalidArgument", c), c)
		return
	}

	var groupChallenge model.GroupChallenge
	groupChallenge.ID = userGroupChallengeRequest.GroupChallengeID
	err = groupChallenge.GetGroupChallenge()
	if err != nil {
		response.Fail(err, localizer.GetMessage("UserGroupChallenge.GetUserGroupChallengeFail", c), c)
		return
	}

	if groupChallenge.ImageID == 0 {
		ok, err := checkFlagFromGroupChallenge(user, userGroupChallengeRequest, groupChallenge)
		if err != nil {
			response.Fail(err, localizer.GetMessage("UserGroupChallenge.CheckFlagFromChallengeFail", c), c)
			return
		}
		if !ok {
			response.Fail(err, localizer.GetMessage("UserGroupChallenge.IncorrectFlag", c), c)
			return
		}
	} else {
		ok, err := checkFlagFromGroupContainer(user, userGroupChallengeRequest, groupChallenge)
		if err != nil {
			response.Fail(err, localizer.GetMessage("UserGroupChallenge.CheckFlagFromContainerFail", c), c)
			return
		}
		if !ok {
			response.Fail(err, localizer.GetMessage("UserGroupChallenge.IncorrectFlag", c), c)
			return
		}
	}

	response.OK(nil, localizer.GetMessage("UserGroupChallenge.CorrectFlag", c), c)
}

func checkFlagFromGroupContainer(user model.User, userGroupChallengeRequest UserGroupChallengeRequest, groupChallenge model.GroupChallenge) (bool, error) {
	userContainers, err := manager.GetUserGroupContainerList(user.Username)
	if err != nil {
		return false, err
	}

	for _, userContainer := range userContainers {
		if userContainer.DockerNodeID == userGroupChallengeRequest.DockerNodeID && userContainer.DockerNodeContainerID == userGroupChallengeRequest.DockerNodeContainerID {
			if userGroupChallengeRequest.Flag == userContainer.Flag {
				err = user.AddGroupChallenge(&groupChallenge)
				if err != nil {
					return false, err
				}
				return true, nil
			}
		}
	}
	return false, nil
}

func checkFlagFromGroupChallenge(user model.User, userGroupChallengeRequest UserGroupChallengeRequest, groupChallenge model.GroupChallenge) (bool, error) {
	if userGroupChallengeRequest.Flag == groupChallenge.Flag {
		err := user.AddGroupChallenge(&groupChallenge)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}
