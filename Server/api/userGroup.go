package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/localizer"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

func GetUserGroupList(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("UserGroup.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroup.UserNotFound", c), c)
		return
	}

	var group model.Group
	err = c.ShouldBindJSON(&group)
	if err != nil {
		response.Fail(err, localizer.GetMessage("UserGroup.InvalidArgument", c), c)
		return
	}

	groupList, err := user.GetGroupList(group)
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroup.GetUserGroupListFail", c), c)
		return
	}

	response.OK(groupList, localizer.GetMessage("UserGroup.GetUserGroupListSuccess", c), c)
}

func GetVisibleUserGroupList(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("UserGroup.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroup.UserNotFound", c), c)
		return
	}

	var group model.Group
	err = c.ShouldBindJSON(&group)
	if err != nil {
		response.Fail(err, localizer.GetMessage("UserGroup.InvalidArgument", c), c)
		return
	}

	groupList, err := user.GetVisibleGroupList(group)
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroup.GetVisibleUserGroupListFail", c), c)
		return
	}

	response.OK(groupList, localizer.GetMessage("UserGroup.GetVisibleUserGroupListSuccess", c), c)
}

func AddUserGroup(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("UserGroup.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroup.UserNotFound", c), c)
		return
	}

	var group model.Group
	err = c.ShouldBindJSON(&group)
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroup.InvalidArgument", c), c)
		return
	}

	err = group.GetGroup()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroup.GetGroupFail", c), c)
		return
	}

	err = user.AddGroup(&group)
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroup.AddUserGroupFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("UserGroup.AddUserGroupSuccess", c), c)
}

func RemoveUserGroup(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("UserGroup.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroup.UserNotFound", c), c)
		return
	}

	var group model.Group
	err = c.ShouldBindJSON(&group)
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroup.InvalidArgument", c), c)
		return
	}

	err = group.GetGroup()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroup.GetGroupFail", c), c)
		return
	}

	err = user.DeleteGroup(&group)
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserGroup.RemoveUserGroupFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("UserGroup.RemoveUserGroupSuccess", c), c)
}
