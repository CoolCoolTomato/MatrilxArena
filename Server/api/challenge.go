package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/localizer"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

func GetChallengeList(c *gin.Context) {
	var challenge model.Challenge
	err := c.ShouldBindJSON(&challenge)
	if err != nil {
		response.Fail(err, localizer.GetMessage("Challenge.InvalidArgument", c), c)
		return
	}

	challengeList, err := model.GetChallengeList(challenge)
	if err != nil {
		response.Fail(err, localizer.GetMessage("Challenge.GetChallengeListFail", c), c)
		return
	}

	response.OK(challengeList, localizer.GetMessage("Challenge.GetChallengeListSuccess", c), c)
}

func GetChallenge(c *gin.Context) {
	var challenge model.Challenge
	err := c.ShouldBindJSON(&challenge)
	if err != nil || challenge.ID == 0 {
		response.Fail(err, localizer.GetMessage("Challenge.InvalidArgument", c), c)
		return
	}

	err = challenge.GetChallenge()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Challenge.GetChallengeFail", c), c)
		return
	}

	response.OK(challenge, localizer.GetMessage("Challenge.GetChallengeSuccess", c), c)
}

func CreateChallenge(c *gin.Context) {
	var challenge model.Challenge
	err := c.ShouldBindJSON(&challenge)
	if err != nil || challenge.Title == "" || challenge.Description == "" || challenge.Flag == "" {
		response.Fail(err, localizer.GetMessage("Challenge.InvalidArgument", c), c)
		return
	}

	err = challenge.CreateChallenge()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Challenge.CreateChallengeFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("Challenge.CreateChallengeSuccess", c), c)
}

func UpdateChallenge(c *gin.Context) {
	var challenge model.Challenge
	err := c.ShouldBindJSON(&challenge)
	if err != nil || challenge.ID == 0 || challenge.Title == "" || challenge.Description == "" || challenge.Flag == "" {
		response.Fail(err, localizer.GetMessage("Challenge.InvalidArgument", c), c)
		return
	}

	err = challenge.UpdateChallenge()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Challenge.UpdateChallengeFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("Challenge.UpdateChallengeSuccess", c), c)
}

func DeleteChallenge(c *gin.Context) {
	var challenge model.Challenge
	err := c.ShouldBindJSON(&challenge)
	if err != nil || challenge.ID == 0 {
		response.Fail(err, localizer.GetMessage("Challenge.InvalidArgument", c), c)
		return
	}

	err = challenge.DeleteChallenge()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Challenge.DeleteChallengeFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("Challenge.DeleteChallengeSuccess", c), c)
}
