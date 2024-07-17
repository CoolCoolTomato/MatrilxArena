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
		response.Fail(err, localizer.GetMessage("User.GetUserListFail", c), c)
		return
	}

	response.OK(userList, localizer.GetMessage("User.GetUserListSuccess", c), c)
}

func GetUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil || user.ID == 0 {
		response.Fail(err, localizer.GetMessage("User.InvalidArgument", c), c)
		return
	}

	err = user.GetUser()
	if err != nil {
		response.Fail(err, localizer.GetMessage("User.GetUserFail", c), c)
		return
	}

	response.OK(user, localizer.GetMessage("User.GetUserSuccess", c), c)
}

func CreateUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil || user.Username == "" || user.Password == "" || user.Email == "" || user.Role == 0 {
		response.Fail(err, localizer.GetMessage("User.InvalidArgument", c), c)
		return
	}
	err = user.SetPassword(user.Password)
	if err != nil {
		response.Fail(err, localizer.GetMessage("User.SetPasswordFail", c), c)
		return
	}
	err = user.CreateUser()
	if err != nil {
		response.Fail(err, localizer.GetMessage("User.CreateUserFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("User.CreateUserSuccess", c), c)
}

func UpdateUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil || user.ID == 0 || user.Username == "" || user.Email == "" || user.Role == 0 {
		response.Fail(err, localizer.GetMessage("User.InvalidArgument", c), c)
		return
	}
	if user.Password != "" {
		err = user.SetPassword(user.Password)
		if err != nil {
			response.Fail(err, localizer.GetMessage("User.SetPasswordFail", c), c)
			return
		}
	}
	err = user.UpdateUser()
	if err != nil {
		response.Fail(err, localizer.GetMessage("User.UpdateUserFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("User.UpdateUserSuccess", c), c)
}

func DeleteUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil || user.ID == 0 {
		response.Fail(err, localizer.GetMessage("User.InvalidArgument", c), c)
		return
	}

	err = user.DeleteUser()
	if err != nil {
		response.Fail(err, localizer.GetMessage("User.DeleteUserFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("User.DeleteUserSuccess", c), c)
}
