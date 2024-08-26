package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/localizer"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

func GetCTFTeamList(c *gin.Context) {
	var ctfTeam model.CTFTeam
	err := c.ShouldBindJSON(&ctfTeam)
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFTeam.InvalidArgument", c), c)
		return
	}

	ctfTeamList, err := model.GetCTFTeamList(ctfTeam)
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFTeam.GetCTFTeamListFail", c), c)
		return
	}

	response.OK(ctfTeamList, localizer.GetMessage("CTFTeam.GetCTFTeamListSuccess", c), c)
}

func GetCTFTeam(c *gin.Context) {
	var ctfTeam model.CTFTeam
	err := c.ShouldBindJSON(&ctfTeam)
	if err != nil || ctfTeam.ID == 0 {
		response.Fail(err, localizer.GetMessage("CTFTeam.InvalidArgument", c), c)
		return
	}

	err = ctfTeam.GetCTFTeam()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFTeam.GetCTFTeamFail", c), c)
		return
	}

	response.OK(ctfTeam, localizer.GetMessage("CTFTeam.GetCTFTeamSuccess", c), c)
}

func CreateCTFTeam(c *gin.Context) {
	var ctfTeam model.CTFTeam
	err := c.ShouldBindJSON(&ctfTeam)
	if err != nil || ctfTeam.Name == "" || ctfTeam.Description == "" {
		response.Fail(err, localizer.GetMessage("CTFTeam.InvalidArgument", c), c)
		return
	}

	err = ctfTeam.CreateCTFTeam()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFTeam.CreateCTFTeamFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("CTFTeam.CreateCTFTeamSuccess", c), c)
}

func UpdateCTFTeam(c *gin.Context) {
	var ctfTeam model.CTFTeam
	err := c.ShouldBindJSON(&ctfTeam)
	if err != nil || ctfTeam.ID == 0 || ctfTeam.Name == "" || ctfTeam.Description == "" {
		response.Fail(err, localizer.GetMessage("CTFTeam.InvalidArgument", c), c)
		return
	}

	err = ctfTeam.UpdateCTFTeam()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFTeam.UpdateCTFTeamFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("CTFTeam.UpdateCTFTeamSuccess", c), c)
}

func DeleteCTFTeam(c *gin.Context) {
	var ctfTeam model.CTFTeam
	err := c.ShouldBindJSON(&ctfTeam)
	if err != nil || ctfTeam.ID == 0 {
		response.Fail(err, localizer.GetMessage("CTFTeam.InvalidArgument", c), c)
		return
	}

	err = ctfTeam.DeleteCTFTeam()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFTeam.DeleteCTFTeamFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("CTFTeam.DeleteCTFTeamSuccess", c), c)
}
