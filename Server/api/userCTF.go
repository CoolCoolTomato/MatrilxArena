package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/localizer"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

func GetUserCTFList(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("UserCTF.InvalidToken", c), c)
		return
	}
    
	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTF.UserNotFound", c), c)
		return
	}

	var ctf model.CTF
	err = c.ShouldBindJSON(&ctf)
	if err != nil {
		response.Fail(err, localizer.GetMessage("UserCTF.InvalidArgument", c), c)
		return
	}

	ctfList, err := user.GetCTFList(ctf)
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTF.GetUserCTFListFail", c), c)
		return
	}

	response.OK(ctfList, localizer.GetMessage("UserCTF.GetUserCTFListSuccess", c), c)
}

func GetVisibleUserCTFList(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("UserCTF.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTF.UserNotFound", c), c)
		return
	}

	var ctf model.CTF
	err = c.ShouldBindJSON(&ctf)
	if err != nil {
		response.Fail(err, localizer.GetMessage("UserCTF.InvalidArgument", c), c)
		return
	}

	ctfList, err := user.GetVisibleCTFList(ctf)
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTF.GetVisibleUserCTFListFail", c), c)
		return
	}

	response.OK(ctfList, localizer.GetMessage("UserCTF.GetVisibleUserCTFListSuccess", c), c)
}

func AddUserCTF(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("UserCTF.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTF.UserNotFound", c), c)
		return
	}

	var ctf model.CTF
	err = c.ShouldBindJSON(&ctf)
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTF.InvalidArgument", c), c)
		return
	}

	err = ctf.GetCTF()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTF.GetCTFFail", c), c)
		return
	}

	err = user.AddCTF(&ctf)
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTF.AddUserCTFFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("UserCTF.AddUserCTFSuccess", c), c)
}

func RemoveUserCTF(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("UserCTF.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTF.UserNotFound", c), c)
		return
	}

	var ctf model.CTF
	err = c.ShouldBindJSON(&ctf)
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTF.InvalidArgument", c), c)
		return
	}

	err = ctf.GetCTF()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTF.GetCTFFail", c), c)
		return
	}

	err = user.DeleteCTF(&ctf)
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserCTF.RemoveUserCTFFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("UserCTF.RemoveUserCTFSuccess", c), c)
}
