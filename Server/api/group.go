package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/localizer"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

type GroupUserRequest struct {
    GroupID uint
    UserID  uint
}

func GetGroupList(c *gin.Context) {
	groupList, err := model.GetGroupList()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Group.GetGroupListFail", c), c)
		return
	}

	response.OK(groupList, localizer.GetMessage("Group.GetGroupListSuccess", c), c)
}

func GetGroup(c *gin.Context) {
	var group model.Group
	err := c.ShouldBindJSON(&group)
	if err != nil || group.ID == 0 {
		response.Fail(err, localizer.GetMessage("Group.InvalidArgument", c), c)
		return
	}

	err = group.GetGroup()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Group.GetGroupFail", c), c)
		return
	}

	response.OK(group, localizer.GetMessage("Group.GetGroupSuccess", c), c)
}

func CreateGroup(c *gin.Context) {
	var group model.Group
	err := c.ShouldBindJSON(&group)
	if err != nil {
		response.Fail(err, localizer.GetMessage("Group.InvalidArgument", c), c)
		return
	}

	err = group.CreateGroup()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Group.CreateGroupFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("Group.CreateGroupSuccess", c), c)
}

func UpdateGroup(c *gin.Context) {
	var group model.Group
	err := c.ShouldBindJSON(&group)
	if err != nil || group.ID == 0 {
		response.Fail(err, localizer.GetMessage("Group.InvalidArgument", c), c)
		return
	}

	err = group.UpdateGroup()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Group.UpdateGroupFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("Group.UpdateGroupSuccess", c), c)
}

func DeleteGroup(c *gin.Context) {
	var group model.Group
	err := c.ShouldBindJSON(&group)
	if err != nil || group.ID == 0 {
		response.Fail(err, localizer.GetMessage("Group.InvalidArgument", c), c)
		return
	}

	err = group.DeleteGroup()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Group.DeleteGroupFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("Group.DeleteGroupSuccess", c), c)
}

func AddGroupUser(c *gin.Context) {
    var groupUserRequest GroupUserRequest
    err := c.ShouldBindJSON(&groupUserRequest)
    if err != nil || groupUserRequest.GroupID == 0 || groupUserRequest.UserID == 0 {
        response.Fail(err, localizer.GetMessage("Group.InvalidArgument", c), c)
		return
    }

    var group model.Group
    group.ID = groupUserRequest.GroupID
    err = group.GetGroup()
    if err != nil {
        response.Fail(err, localizer.GetMessage("Group.GetGroupFail", c), c)
		return
    }

    var user model.User
    user.ID = groupUserRequest.UserID
    err = user.GetUser()
    if err != nil {
        response.Fail(err, localizer.GetMessage("Group.GetUserFail", c), c)
		return
    }

    err = group.AddUser(&user)
    if err != nil {
        response.Fail(err, localizer.GetMessage("Group.AddUserFail", c), c)
		return
    }

    response.OK(nil, localizer.GetMessage("Group.AddUserSuccess", c), c)
}

func RemoveGroupUser(c *gin.Context) {
    var groupUserRequest GroupUserRequest
    err := c.ShouldBindJSON(&groupUserRequest)
    if err != nil || groupUserRequest.GroupID == 0 || groupUserRequest.UserID == 0 {
        response.Fail(err, localizer.GetMessage("Group.InvalidArgument", c), c)
		return
    }

    var group model.Group
    group.ID = groupUserRequest.GroupID
    err = group.GetGroup()
    if err != nil {
        response.Fail(err, localizer.GetMessage("Group.GetGroupFail", c), c)
		return
    }

    var user model.User
    user.ID = groupUserRequest.UserID
    err = user.GetUser()
    if err != nil {
        response.Fail(err, localizer.GetMessage("Group.GetUserFail", c), c)
		return
    }

    err = group.DeleteUser(&user)
    if err != nil {
        response.Fail(err, localizer.GetMessage("Group.RemoveUserFail", c), c)
		return
    }

    response.OK(nil, localizer.GetMessage("Group.RemoveUserSuccess", c), c)
}
