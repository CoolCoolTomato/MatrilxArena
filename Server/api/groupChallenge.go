package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/localizer"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

func GetGroupChallengeList(c *gin.Context) {
	var groupChallenge model.GroupChallenge
	err := c.ShouldBindJSON(&groupChallenge)
	if err != nil {
		response.Fail(err, localizer.GetMessage("GroupChallenge.InvalidArgument", c), c)
		return
	}

	groupChallengeList, err := model.GetGroupChallengeList(groupChallenge)
	if err != nil {
		response.Fail(err, localizer.GetMessage("GroupChallenge.GetGroupChallengeListFail", c), c)
		return
	}
	response.OK(groupChallengeList, localizer.GetMessage("GroupChallenge.GetGroupChallengeListSuccess", c), c)
}

func GetGroupChallenge(c *gin.Context) {
	var groupChallenge model.GroupChallenge
	err := c.ShouldBindJSON(&groupChallenge)
	if err != nil || groupChallenge.ID == 0 {
		response.Fail(err, localizer.GetMessage("GroupChallenge.InvalidArgument", c), c)
		return
	}

	err = groupChallenge.GetGroupChallenge()
	if err != nil {
		response.Fail(err, localizer.GetMessage("GroupChallenge.GetGroupChallengeFail", c), c)
		return
	}

	response.OK(groupChallenge, localizer.GetMessage("GroupChallenge.GetGroupChallengeSuccess", c), c)
}

func CreateGroupChallenge(c *gin.Context) {
	var groupChallenge model.GroupChallenge
	err := c.ShouldBindJSON(&groupChallenge)
	if err != nil || groupChallenge.Title == "" || groupChallenge.Description == "" || groupChallenge.Flag == "" {
		response.Fail(err, localizer.GetMessage("GroupChallenge.InvalidArgument", c), c)
		return
	}

	err = groupChallenge.CreateGroupChallenge()
	if err != nil {
		response.Fail(err, localizer.GetMessage("GroupChallenge.CreateGroupChallengeFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("GroupChallenge.CreateGroupChallengeSuccess", c), c)
}

func UpdateGroupChallenge(c *gin.Context) {
	var groupChallenge model.GroupChallenge
	err := c.ShouldBindJSON(&groupChallenge)
	if err != nil || groupChallenge.ID == 0 || groupChallenge.Title == "" || groupChallenge.Description == "" || groupChallenge.Flag == "" {
		response.Fail(err, localizer.GetMessage("GroupChallenge.InvalidArgument", c), c)
		return
	}

	err = groupChallenge.UpdateGroupChallenge()
	if err != nil {
		response.Fail(err, localizer.GetMessage("GroupChallenge.UpdateGroupChallengeFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("GroupChallenge.UpdateGroupChallengeSuccess", c), c)
}

func DeleteGroupChallenge(c *gin.Context) {
	var groupChallenge model.GroupChallenge
	err := c.ShouldBindJSON(&groupChallenge)
	if err != nil || groupChallenge.ID == 0 {
		response.Fail(err, localizer.GetMessage("GroupChallenge.InvalidArgument", c), c)
		return
	}

	err = groupChallenge.DeleteGroupChallenge()
	if err != nil {
		response.Fail(err, localizer.GetMessage("GroupChallenge.DeleteGroupChallengeFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("GroupChallenge.DeleteGroupChallengeSuccess", c), c)
}
