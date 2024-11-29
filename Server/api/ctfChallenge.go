package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/challenge"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/localizer"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

func GetCTFChallengeList(c *gin.Context) {
	var ctfChallenge model.CTFChallenge
	err := c.ShouldBindJSON(&ctfChallenge)
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFChallenge.InvalidArgument", c), c)
		return
	}

	ctfChallengeList, err := model.GetCTFChallengeList(ctfChallenge)
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFChallenge.GetCTFChallengeListFail", c), c)
		return
	}

	ctfChallengeList = challenge.CalculateCTFChallengeListScore(ctfChallengeList)

	response.OK(ctfChallengeList, localizer.GetMessage("CTFChallenge.GetCTFChallengeListSuccess", c), c)
}

func GetCTFChallenge(c *gin.Context) {
	var ctfChallenge model.CTFChallenge
	err := c.ShouldBindJSON(&ctfChallenge)
	if err != nil || ctfChallenge.ID == 0 {
		response.Fail(err, localizer.GetMessage("CTFChallenge.InvalidArgument", c), c)
		return
	}

	err = ctfChallenge.GetCTFChallenge()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFChallenge.GetCTFChallengeFail", c), c)
		return
	}

	ctfChallenge = challenge.CalculateCTFChallengeScore(ctfChallenge)

	response.OK(ctfChallenge, localizer.GetMessage("CTFChallenge.GetCTFChallengeSuccess", c), c)
}

func CreateCTFChallenge(c *gin.Context) {
	var ctfChallenge model.CTFChallenge
	err := c.ShouldBindJSON(&ctfChallenge)
	if err != nil || ctfChallenge.Title == "" || ctfChallenge.Description == "" || ctfChallenge.Flag == "" || ctfChallenge.Score == 0 {
		response.Fail(err, localizer.GetMessage("CTFChallenge.InvalidArgument", c), c)
		return
	}

	err = ctfChallenge.CreateCTFChallenge()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFChallenge.CreateCTFChallengeFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("CTFChallenge.CreateCTFChallengeSuccess", c), c)
}

func UpdateCTFChallenge(c *gin.Context) {
	var ctfChallenge model.CTFChallenge
	err := c.ShouldBindJSON(&ctfChallenge)
	if err != nil || ctfChallenge.ID == 0 || ctfChallenge.Title == "" || ctfChallenge.Description == "" || ctfChallenge.Flag == "" || ctfChallenge.Score == 0 {
		response.Fail(err, localizer.GetMessage("CTFChallenge.InvalidArgument", c), c)
		return
	}

	err = ctfChallenge.UpdateCTFChallenge()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFChallenge.UpdateCTFChallengeFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("CTFChallenge.UpdateCTFChallengeSuccess", c), c)
}

func DeleteCTFChallenge(c *gin.Context) {
	var ctfChallenge model.CTFChallenge
	err := c.ShouldBindJSON(&ctfChallenge)
	if err != nil || ctfChallenge.ID == 0 {
		response.Fail(err, localizer.GetMessage("CTFChallenge.InvalidArgument", c), c)
		return
	}

	err = ctfChallenge.DeleteCTFChallenge()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFChallenge.DeleteCTFChallengeFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("CTFChallenge.DeleteCTFChallengeSuccess", c), c)
}
