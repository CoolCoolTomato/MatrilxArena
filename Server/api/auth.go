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
		response.Fail(err, localizer.GetMessage("InvalidArgument", c), c)
		return
	}

	dbUser := loginUser
	err = dbUser.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(err, localizer.GetMessage("UserNotFound", c), c)
		return
	}

	err = dbUser.CheckPassword(loginUser.Password)
	if err != nil {
		response.Fail(err, localizer.GetMessage("InvalidUsernameOrPassword", c), c)
		return
	}

	token, err := jwt.GenerateJWT(dbUser.Username)
	if err != nil {
		response.Fail(err, localizer.GetMessage("FailedToGenerateToken", c), c)
		return
	}

	response.OK(token, localizer.GetMessage("LoginSuccess", c), c)
}

func GetUserByAuth(c *gin.Context) {
	username, exists := c.Get("Username")
	if !exists {
		response.Fail(nil, localizer.GetMessage("InvalidToken", c), c)
		return
	}

	var user model.User
	user.Username = username.(string)
	err := user.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(nil, localizer.GetMessage("UserNotFound", c), c)
		return
	}

	response.OK(user, localizer.GetMessage("GetUserByAuthSuccess", c), c)
}
