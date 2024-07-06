package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/localizer"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	userList, err := model.GetUserList()
	if err != nil {
		response.Fail(err, localizer.GetMessage("GetUserListFail", c), c)
		return
	}

	response.OK(userList, localizer.GetMessage("GetUserListSuccess", c), c)
}

func GetUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil || user.ID == 0 {
		response.Fail(err, localizer.GetMessage("InvalidArgument", c), c)
		return
	}

	err = user.GetUser()
	if err != nil {
		response.Fail(err, localizer.GetMessage("GetUserFail", c), c)
		return
	}

	response.OK(user, localizer.GetMessage("GetUserSuccess", c), c)
}

func CreateUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil || user.Username == "" || user.Password == "" || user.Email == "" || user.Role == 0 {
		response.Fail(err, localizer.GetMessage("InvalidArgument", c), c)
		return
	}
	err = user.SetPassword(user.Password)
	if err != nil {
		response.Fail(err, localizer.GetMessage("SetPasswordFail", c), c)
		return
	}
	err = user.CreateUser()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CreateUserFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("CreateUserSuccess", c), c)
}

func UpdateUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil || user.ID == 0 || user.Username == "" || user.Email == "" || user.Role == 0 {
		response.Fail(err, localizer.GetMessage("InvalidArgument", c), c)
		return
	}
	if user.Password != "" {
		err = user.SetPassword(user.Password)
		if err != nil {
			response.Fail(err, localizer.GetMessage("SetPasswordFail", c), c)
			return
		}
	}
	err = user.UpdateUser()
	if err != nil {
		response.Fail(err, localizer.GetMessage("UpdateUserFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("UpdateUserSuccess", c), c)
}

func DeleteUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil || user.ID == 0 {
		response.Fail(err, localizer.GetMessage("InvalidArgument", c), c)
		return
	}

	err = user.DeleteUser()
	if err != nil {
		response.Fail(err, localizer.GetMessage("DeleteUserFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("DeleteUserSuccess", c), c)
}
