package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/localizer"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

func GetCTFList(c *gin.Context) {
	var ctf model.CTF
	err := c.ShouldBindJSON(&ctf)
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTF.InvalidArgument", c), c)
		return
	}

	ctfList, err := model.GetCTFList(ctf)
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTF.GetCTFListFail", c), c)
		return
	}

	response.OK(ctfList, localizer.GetMessage("CTF.GetCTFListSuccess", c), c)
}

func GetCTF(c *gin.Context) {
	var ctf model.CTF
	err := c.ShouldBindJSON(&ctf)
	if err != nil || ctf.ID == 0 {
		response.Fail(err, localizer.GetMessage("CTF.InvalidArgument", c), c)
		return
	}

	err = ctf.GetCTF()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTF.GetCTFFail", c), c)
		return
	}

	response.OK(ctf, localizer.GetMessage("CTF.GetCTFSuccess", c), c)
}

func CreateCTF(c *gin.Context) {
	var ctf model.CTF
	err := c.ShouldBindJSON(&ctf)
	if err != nil || ctf.Name == "" || ctf.Description == "" {
		response.Fail(err, localizer.GetMessage("CTF.InvalidArgument", c), c)
		return
	}

	err = ctf.CreateCTF()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTF.CreateCTFFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("CTF.CreateCTFSuccess", c), c)
}

func UpdateCTF(c *gin.Context) {
	var ctf model.CTF
	err := c.ShouldBindJSON(&ctf)
	if err != nil || ctf.ID == 0 || ctf.Name == "" || ctf.Description == "" {
		response.Fail(err, localizer.GetMessage("CTF.InvalidArgument", c), c)
		return
	}

	err = ctf.UpdateCTF()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTF.UpdateCTFFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("CTF.UpdateCTFSuccess", c), c)
}

func DeleteCTF(c *gin.Context) {
	var ctf model.CTF
	err := c.ShouldBindJSON(&ctf)
	if err != nil || ctf.ID == 0 {
		response.Fail(err, localizer.GetMessage("CTF.InvalidArgument", c), c)
		return
	}

	err = ctf.DeleteCTF()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CTF.DeleteCTFFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("CTF.DeleteCTFSuccess", c), c)
}
