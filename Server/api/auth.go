package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/common/response"
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginUser model.User
	err := c.ShouldBindJSON(&loginUser)
	if err != nil || (loginUser.Username == "" && loginUser.Password == "") {
		response.Fail(err, "Invalid argument", c)
		return
	}

	dbUser := loginUser
	err = dbUser.GetUserByUsernameOrEmail()
	if err != nil {
		response.Fail(err, "User not found", c)
		return
	}

	err = dbUser.CheckPassword(loginUser.Password)
	if err != nil {
		response.Fail(err, "Invalid username or password", c)
		return
	}

	token, err := utils.GenerateJWT(dbUser.Username)
	if err != nil {
		response.Fail(err, "Failed to generate token", c)
		return
	}

	response.OK(token, "Login success", c)
}
