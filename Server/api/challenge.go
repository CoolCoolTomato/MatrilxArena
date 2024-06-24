package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

func GetChallengeList(c *gin.Context) {
	challengeList, err := model.GetChallengeList()
	if err != nil {
		response.Fail(err, "Get challenge list fail", c)
		return
	}
	response.OK(challengeList, "Get challenge list success", c)
}

func GetChallenge(c *gin.Context) {
	var challenge model.Challenge
	err := c.ShouldBindJSON(&challenge)
	if err != nil || challenge.ID == 0 {
		response.Fail(err, "Invalid argument", c)
		return
	}

	err = challenge.GetChallenge()
	if err != nil {
		response.Fail(err, "Get challenge fail", c)
		return
	}

	response.OK(challenge, "Get challenge success", c)
}

func CreateChallenge(c *gin.Context) {
	var challenge model.Challenge
	err := c.ShouldBindJSON(&challenge)
	if err != nil || challenge.Title == "" || challenge.Description == "" || challenge.ImageID == 0 || challenge.Flag == "" {
		response.Fail(err, "Invalid argument", c)
		return
	}

	err = challenge.CreateChallenge()
	if err != nil {
		response.Fail(err, "Create challenge fail", c)
		return
	}

	response.OK(nil, "Create challenge success", c)
}

func UpdateChallenge(c *gin.Context) {
	var challenge model.Challenge
	err := c.ShouldBindJSON(&challenge)
	if err != nil || challenge.ID == 0 || challenge.Title == "" || challenge.Description == "" || challenge.ImageID == 0 || challenge.Flag == "" {
		response.Fail(err, "Invalid argument", c)
		return
	}

	err = challenge.UpdateChallenge()
	if err != nil {
		response.Fail(err, "Update challenge fail", c)
		return
	}

	response.OK(nil, "Update challenge success", c)
}

func DeleteChallenge(c *gin.Context) {
	var challenge model.Challenge
	err := c.ShouldBindJSON(&challenge)
	if err != nil || challenge.ID == 0 {
		response.Fail(err, "Invalid argument", c)
		return
	}

	err = challenge.DeleteChallenge()
	if err != nil {
		response.Fail(err, "Delete challenge fail", c)
		return
	}

	response.OK(nil, "Delete challenge success", c)
}
