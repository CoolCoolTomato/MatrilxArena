package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/common/response"
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
)

func GetUserList(c *gin.Context) {
	userList, err := model.GetUserList()
	if err != nil {
		response.Fail(err, "Get user list fail", c)
		return
	}

	response.OK(userList, "Get user list success", c)
}

func GetUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil || user.ID == 0 {
		response.Fail(err, "Invalid argument", c)
		return
	}
	
	err = user.GetUser()
	if err != nil {
		response.Fail(err, "Get user fail", c)
		return
	}
	
	response.OK(user, "Get user success", c)
}

func CreateUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil || user.Username == "" || user.Password == "" || user.Email == "" {
		response.Fail(err, "Invalid argument", c)
		return
	}
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    user.Password = string(hashPassword)
	err = user.CreateUser()
	if err != nil {
		response.Fail(err, "Create user fail", c)
		return
	}
	
	response.OK(nil, "Create user success", c)
}

func UpdateUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil || user.ID == 0 || user.Username == "" || user.Password == "" || user.Email == "" {
		response.Fail(err, "Invalid argument", c)
		return
	}
    hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    user.Password = string(hashPassword)
	err = user.UpdateUser()
	if err != nil {
		response.Fail(err, "Update user fail", c)
		return
	}
	
	response.OK(nil, "Update user success", c)
}

func DeleteUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil || user.ID == 0 {
		response.Fail(err, "Invalid argument", c)
		return
	}
	
	err = user.DeleteUser()
	if err != nil {
		response.Fail(err, "Delete user fail", c)
		return
	}
	
	response.OK(nil, "Delete user success", c)
}
