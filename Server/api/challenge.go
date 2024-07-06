package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/localizer"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

func GetChallengeList(c *gin.Context) {
	challengeList, err := model.GetChallengeList()
	if err != nil {
		response.Fail(err, localizer.GetMessage("GetChallengeListFail", c), c)
		return
	}
	response.OK(challengeList, localizer.GetMessage("GetChallengeListSuccess", c), c)
}

func GetChallengeListByQuery(c *gin.Context) {
	var challenge model.Challenge
	err := c.ShouldBindJSON(&challenge)
	if err != nil {
		response.Fail(err, localizer.GetMessage("InvalidArgument", c), c)
		return
	}

	challengeList, err := model.GetChallengeListByQuery(challenge)
	if err != nil {
		response.Fail(err, localizer.GetMessage("GetChallengeListFail", c), c)
		return
	}
	response.OK(challengeList, localizer.GetMessage("GetChallengeListSuccess", c), c)
}

func GetChallenge(c *gin.Context) {
	var challenge model.Challenge
	err := c.ShouldBindJSON(&challenge)
	if err != nil || challenge.ID == 0 {
		response.Fail(err, localizer.GetMessage("InvalidArgument", c), c)
		return
	}

	err = challenge.GetChallenge()
	if err != nil {
		response.Fail(err, localizer.GetMessage("GetChallengeFail", c), c)
		return
	}

	response.OK(challenge, localizer.GetMessage("GetChallengeSuccess", c), c)
}

func CreateChallenge(c *gin.Context) {
	var challenge model.Challenge
	err := c.ShouldBindJSON(&challenge)
	if err != nil || challenge.Title == "" || challenge.Description == "" || challenge.Flag == "" {
		response.Fail(err, localizer.GetMessage("InvalidArgument", c), c)
		return
	}

	err = challenge.CreateChallenge()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CreateChallengeFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("CreateChallengeSuccess", c), c)
}

func UpdateChallenge(c *gin.Context) {
	var challenge model.Challenge
	err := c.ShouldBindJSON(&challenge)
	if err != nil || challenge.ID == 0 || challenge.Title == "" || challenge.Description == "" || challenge.Flag == "" {
		response.Fail(err, localizer.GetMessage("InvalidArgument", c), c)
		return
	}

	err = challenge.UpdateChallenge()
	if err != nil {
		response.Fail(err, localizer.GetMessage("UpdateChallengeFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("UpdateChallengeSuccess", c), c)
}

func DeleteChallenge(c *gin.Context) {
	var challenge model.Challenge
	err := c.ShouldBindJSON(&challenge)
	if err != nil || challenge.ID == 0 {
		response.Fail(err, localizer.GetMessage("InvalidArgument", c), c)
		return
	}

	err = challenge.DeleteChallenge()
	if err != nil {
		response.Fail(err, localizer.GetMessage("DeleteChallengeFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("DeleteChallengeSuccess", c), c)
}
