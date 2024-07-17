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
		response.Fail(err, localizer.GetMessage("Image.GetImageListFail", c), c)
		return
	}

	response.OK(imageList, localizer.GetMessage("Image.GetImageListSuccess", c), c)
}

func GetImage(c *gin.Context) {
	var image model.Image
	err := c.ShouldBindJSON(&image)
	if err != nil || image.ID == 0 {
		response.Fail(err, localizer.GetMessage("Image.InvalidArgument", c), c)
		return
	}

	err = image.GetImage()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Image.GetImageFail", c), c)
		return
	}

	response.OK(image, localizer.GetMessage("Image.GetImageSuccess", c), c)
}

func CreateImage(c *gin.Context) {
	var image model.Image
	err := c.ShouldBindJSON(&image)
	if err != nil || image.Remark == "" || image.RepoTags == "" {
		response.Fail(err, localizer.GetMessage("Image.InvalidArgument", c), c)
		return
	}

	err = image.CreateImage()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Image.CreateImageFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("Image.CreateImageSuccess", c), c)
}

func UpdateImage(c *gin.Context) {
	var image model.Image
	err := c.ShouldBindJSON(&image)
	if err != nil || image.ID == 0 || image.Remark == "" || image.RepoTags == "" {
		response.Fail(err, localizer.GetMessage("Image.InvalidArgument", c), c)
		return
	}

	err = image.UpdateImage()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Image.UpdateImageFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("Image.UpdateImageSuccess", c), c)
}

func DeleteImage(c *gin.Context) {
	var image model.Image
	err := c.ShouldBindJSON(&image)
	if err != nil || image.ID == 0 {
		response.Fail(err, localizer.GetMessage("Image.InvalidArgument", c), c)
		return
	}

	err = image.DeleteImage()
	if err != nil {
		response.Fail(err, localizer.GetMessage("Image.DeleteImageFail", c), c)
		return
	}

	response.OK(nil, localizer.GetMessage("Image.DeleteImageSuccess", c), c)
}
