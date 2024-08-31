package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/localizer"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

type CTFUserRequest struct {
	CTFID  uint
	UserID   uint
	Username string
}

func GetCTFUserList(c *gin.Context) {
	var ctfUserRequest CTFUserRequest
	err := c.ShouldBindJSON(&ctfUserRequest)
	if err != nil || ctfUserRequest.CTFID == 0 {
		response.Fail(err, localizer.GetMessage("CTFUser.InvalidArgument", c), c)
		return
	}

	var ctf model.CTF
	ctf.ID = ctfUserRequest.CTFID
	err = ctf.GetCTF()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFUser.GetCTFFail", c), c)
		return
	}

	var user model.User
	user.Username = ctfUserRequest.Username
	userList, err := ctf.GetUserList(user)
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFUser.GetUserListFail", c), c)
		return
	}

	response.OK(userList, localizer.GetMessage("CTFUser.GetUserListSuccess", c), c)
}

func AddCTFUser(c *gin.Context) {
	var ctfUserRequest CTFUserRequest
	err := c.ShouldBindJSON(&ctfUserRequest)
	if err != nil || ctfUserRequest.CTFID == 0 || ctfUserRequest.UserID == 0 {
		response.Fail(err, localizer.GetMessage("CTFUser.InvalidArgument", c), c)
		return
	}

	var ctf model.CTF
	ctf.ID = ctfUserRequest.CTFID
	err = ctf.GetCTF()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFUser.GetCTFFail", c), c)
		return
	}

	var user model.User
	user.ID = ctfUserRequest.UserID
	err = user.GetUser()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFUser.GetUserFail", c), c)
		return
	}

	err = ctf.AddUser(&user)
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFUser.AddUserFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("CTFUser.AddUserSuccess", c), c)
}

func RemoveCTFUser(c *gin.Context) {
	var ctfUserRequest CTFUserRequest
	err := c.ShouldBindJSON(&ctfUserRequest)
	if err != nil || ctfUserRequest.CTFID == 0 || ctfUserRequest.UserID == 0 {
		response.Fail(err, localizer.GetMessage("CTFUser.InvalidArgument", c), c)
		return
	}

	var ctf model.CTF
	ctf.ID = ctfUserRequest.CTFID
	err = ctf.GetCTF()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFUser.GetCTFFail", c), c)
		return
	}

	var user model.User
	user.ID = ctfUserRequest.UserID
	err = user.GetUser()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFUser.GetUserFail", c), c)
		return
	}

	err = ctf.DeleteUser(&user)
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTFUser.RemoveUserFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("CTFUser.RemoveUserSuccess", c), c)
}
