package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

func GetChallengeClassList(c *gin.Context) {
	challengeClassList, err := model.GetChallengeClassList()
	if err != nil {
		response.Fail(err, "Get challengeClass list fail", c)
		return
	}

	response.OK(challengeClassList, "Get challengeClass list success", c)
}

func GetChallengeClass(c *gin.Context) {
	var challengeClass model.ChallengeClass
	err := c.ShouldBindJSON(&challengeClass)
	if err != nil || challengeClass.ID == 0 {
		response.Fail(err, "Invalid argument", c)
		return
	}

	err = challengeClass.GetChallengeClass()
	if err != nil {
		response.Fail(err, "Get challengeClass fail", c)
		return
	}

	response.OK(challengeClass, "Get challengeClass success", c)
}

func CreateChallengeClass(c *gin.Context) {
	var challengeClass model.ChallengeClass
	err := c.ShouldBindJSON(&challengeClass)
	if err != nil {
		response.Fail(err, "Invalid argument", c)
		return
	}

	err = challengeClass.CreateChallengeClass()
	if err != nil {
		response.Fail(err, "Create challengeClass fail", c)
		return
	}

	response.OK(nil, "Create challengeClass success", c)
}

func UpdateChallengeClass(c *gin.Context) {
	var challengeClass model.ChallengeClass
	err := c.ShouldBindJSON(&challengeClass)
	if err != nil || challengeClass.ID == 0 {
		response.Fail(err, "Invalid argument", c)
		return
	}

	err = challengeClass.UpdateChallengeClass()
	if err != nil {
		response.Fail(err, "Update challengeClass fail", c)
		return
	}

	response.OK(nil, "Update challengeClass success", c)
}

func DeleteChallengeClass(c *gin.Context) {
	var challengeClass model.ChallengeClass
	err := c.ShouldBindJSON(&challengeClass)
	if err != nil || challengeClass.ID == 0 {
		response.Fail(err, "Invalid argument", c)
		return
	}

	err = challengeClass.DeleteChallengeClass()
	if err != nil {
		response.Fail(err, "Delete challengeClass fail", c)
		return
	}

	response.OK(nil, "Delete challengeClass success", c)
}
