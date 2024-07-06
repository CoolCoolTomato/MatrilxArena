package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/localizer"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

func GetImageList(c *gin.Context) {
	imageList, err := model.GetImageList()
	if err != nil {
		response.Fail(err, localizer.GetMessage("GetImageListFail", c), c)
		return
	}

	response.OK(imageList, localizer.GetMessage("GetImageListSuccess", c), c)
}

func GetImage(c *gin.Context) {
	var image model.Image
	err := c.ShouldBindJSON(&image)
	if err != nil || image.ID == 0 {
		response.Fail(err, localizer.GetMessage("InvalidArgument", c), c)
		return
	}

	err = image.GetImage()
	if err != nil {
		response.Fail(err, localizer.GetMessage("GetImageFail", c), c)
		return
	}

	response.OK(image, localizer.GetMessage("GetImageSuccess", c), c)
}

func CreateImage(c *gin.Context) {
	var image model.Image
	err := c.ShouldBindJSON(&image)
	if err != nil {
		response.Fail(err, localizer.GetMessage("InvalidArgument", c), c)
		return
	}

	err = image.CreateImage()
	if err != nil {
		response.Fail(err, localizer.GetMessage("CreateImageFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("CreateImageSuccess", c), c)
}

func UpdateImage(c *gin.Context) {
	var image model.Image
	err := c.ShouldBindJSON(&image)
	if err != nil || image.ID == 0 {
		response.Fail(err, localizer.GetMessage("InvalidArgument", c), c)
		return
	}

	err = image.UpdateImage()
	if err != nil {
		response.Fail(err, localizer.GetMessage("UpdateImageFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("UpdateImageSuccess", c), c)
}

func DeleteImage(c *gin.Context) {
	var image model.Image
	err := c.ShouldBindJSON(&image)
	if err != nil || image.ID == 0 {
		response.Fail(err, localizer.GetMessage("InvalidArgument", c), c)
		return
	}

	err = image.DeleteImage()
	if err != nil {
		response.Fail(err, localizer.GetMessage("DeleteImageFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("DeleteImageSuccess", c), c)
}
