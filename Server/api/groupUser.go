package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/localizer"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

type GroupUserRequest struct {
	GroupID  uint
	UserID   uint
	Username string
}

func GetGroupUserList(c *gin.Context) {
	var groupUserRequest GroupUserRequest
	err := c.ShouldBindJSON(&groupUserRequest)
	if err != nil || groupUserRequest.GroupID == 0 {
		response.Fail(err, localizer.GetMessage("GroupUser.InvalidArgument", c), c)
		return
	}

	var group model.Group
	group.ID = groupUserRequest.GroupID
	err = group.GetGroup()
	if err != nil {
		response.Fail(err, localizer.GetMessage("GroupUser.GetGroupFail", c), c)
		return
	}

	var user model.User
	user.Username = groupUserRequest.Username
	userList, err := group.GetUserList(user)
	if err != nil {
		response.Fail(err, localizer.GetMessage("GroupUser.GetUserListFail", c), c)
		return
	}

	response.OK(userList, localizer.GetMessage("GroupUser.GetUserListSuccess", c), c)
}

func AddGroupUser(c *gin.Context) {
	var groupUserRequest GroupUserRequest
	err := c.ShouldBindJSON(&groupUserRequest)
	if err != nil || groupUserRequest.GroupID == 0 || groupUserRequest.UserID == 0 {
		response.Fail(err, localizer.GetMessage("GroupUser.InvalidArgument", c), c)
		return
	}

	var group model.Group
	group.ID = groupUserRequest.GroupID
	err = group.GetGroup()
	if err != nil {
		response.Fail(err, localizer.GetMessage("GroupUser.GetGroupFail", c), c)
		return
	}

	var user model.User
	user.ID = groupUserRequest.UserID
	err = user.GetUser()
	if err != nil {
		response.Fail(err, localizer.GetMessage("GroupUser.GetUserFail", c), c)
		return
	}

	err = group.AddUser(&user)
	if err != nil {
		response.Fail(err, localizer.GetMessage("GroupUser.AddUserFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("GroupUser.AddUserSuccess", c), c)
}

func RemoveGroupUser(c *gin.Context) {
	var groupUserRequest GroupUserRequest
	err := c.ShouldBindJSON(&groupUserRequest)
	if err != nil || groupUserRequest.GroupID == 0 || groupUserRequest.UserID == 0 {
		response.Fail(err, localizer.GetMessage("GroupUser.InvalidArgument", c), c)
		return
	}

	var group model.Group
	group.ID = groupUserRequest.GroupID
	err = group.GetGroup()
	if err != nil {
		response.Fail(err, localizer.GetMessage("GroupUser.GetGroupFail", c), c)
		return
	}

	var user model.User
	user.ID = groupUserRequest.UserID
	err = user.GetUser()
	if err != nil {
		response.Fail(err, localizer.GetMessage("GroupUser.GetUserFail", c), c)
		return
	}

	err = group.DeleteUser(&user)
	if err != nil {
		response.Fail(err, localizer.GetMessage("GroupUser.RemoveUserFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("GroupUser.RemoveUserSuccess", c), c)
}
