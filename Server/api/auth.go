package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/jwt"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/localizer"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginUser model.User
	err := c.ShouldBindJSON(&loginUser)
	if err != nil || (loginUser.Username == "" && loginUser.Password == "") {
		response.Fail(err, localizer.GetMessage("Auth.InvalidArgument", c), c)
		return
	}

	dbUser := loginUser
	err = dbUser.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Auth.UserNotFound", c), c)
		return
	}

	err = dbUser.CheckPassword(loginUser.Password)
	if err != nil {
		response.Fail(err, localizer.GetMessage("Auth.InvalidUsernameOrPassword", c), c)
		return
	}

	token, err := jwt.GenerateJWT(dbUser.Username)
	if err != nil {
		response.Fail(err, localizer.GetMessage("Auth.FailedToGenerateToken", c), c)
		return
	}

	response.OK(token, localizer.GetMessage("Auth.LoginSuccess", c), c)
}

func GetUserByAuth(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("Auth.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("Auth.UserNotFound", c), c)
		return
	}

	response.OK(user, localizer.GetMessage("Auth.GetUserByAuthSuccess", c), c)
}

func ResetPassword(c *gin.Context) {
    username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("Auth.InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("Auth.UserNotFound", c), c)
		return
	}

    var userRequest model.User
    err = c.ShouldBindJSON(&userRequest)
    if err != nil || userRequest.Password == "" {
		response.Fail(err, localizer.GetMessage("Auth.InvalidArgument", c), c)
		return
	}

    err = user.SetPassword(userRequest.Password)
    if err != nil {
        response.Fail(err, localizer.GetMessage("Auth.SetPasswordFail", c), c)
		return
    }

    err = user.UpdateUser()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Auth.UpdateUserFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("Auth.ResetPasswordSuccess", c), c)
}
