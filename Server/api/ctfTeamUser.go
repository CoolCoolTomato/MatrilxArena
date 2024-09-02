package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/localizer"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

type CTFTeamUserRequest struct {
	CTFTeamID  uint
	UserID   uint
	Username string
}

func GetCTFTeamUserList(c *gin.Context) {
	var ctfTeamUserRequest CTFTeamUserRequest
	err := c.ShouldBindJSON(&ctfTeamUserRequest)
	if err != nil || ctfTeamUserRequest.CTFTeamID == 0 {
		response.Fail(err, localizer.GetMessage("CTFTeamUser.InvalidArgument", c), c)
		return
	}

	var ctfTeam model.CTFTeam
	ctfTeam.ID = ctfTeamUserRequest.CTFTeamID
	err = ctfTeam.GetCTFTeam()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFTeamUser.GetCTFTeamFail", c), c)
		return
	}

	var user model.User
	user.Username = ctfTeamUserRequest.Username
	userList, err := ctfTeam.GetUserList(user)
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFTeamUser.GetUserListFail", c), c)
		return
	}

	response.OK(userList, localizer.GetMessage("CTFTeamUser.GetUserListSuccess", c), c)
}

func AddCTFTeamUser(c *gin.Context) {
	var ctfTeamUserRequest CTFTeamUserRequest
	err := c.ShouldBindJSON(&ctfTeamUserRequest)
	if err != nil || ctfTeamUserRequest.CTFTeamID == 0 || ctfTeamUserRequest.UserID == 0 {
		response.Fail(err, localizer.GetMessage("CTFTeamUser.InvalidArgument", c), c)
		return
	}

	var ctfTeam model.CTFTeam
	ctfTeam.ID = ctfTeamUserRequest.CTFTeamID
	err = ctfTeam.GetCTFTeam()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFTeamUser.GetCTFTeamFail", c), c)
		return
	}

	var user model.User
	user.ID = ctfTeamUserRequest.UserID
	err = user.GetUser()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFTeamUser.GetUserFail", c), c)
		return
	}

	err = ctfTeam.AddUser(&user)
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFTeamUser.AddUserFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("CTFTeamUser.AddUserSuccess", c), c)
}

func RemoveCTFTeamUser(c *gin.Context) {
	var ctfTeamUserRequest CTFTeamUserRequest
	err := c.ShouldBindJSON(&ctfTeamUserRequest)
	if err != nil || ctfTeamUserRequest.CTFTeamID == 0 || ctfTeamUserRequest.UserID == 0 {
		response.Fail(err, localizer.GetMessage("CTFTeamUser.InvalidArgument", c), c)
		return
	}

	var ctfTeam model.CTFTeam
	ctfTeam.ID = ctfTeamUserRequest.CTFTeamID
	err = ctfTeam.GetCTFTeam()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFTeamUser.GetCTFTeamFail", c), c)
		return
	}

	var user model.User
	user.ID = ctfTeamUserRequest.UserID
	err = user.GetUser()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFTeamUser.GetUserFail", c), c)
		return
	}

	err = ctfTeam.DeleteUser(&user)
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFTeamUser.RemoveUserFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("CTFTeamUser.RemoveUserSuccess", c), c)
}
