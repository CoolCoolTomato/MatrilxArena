package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/common/response"
	"github.com/CoolCoolTomato/MatrilxArena/Server/docker"
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/gin-gonic/gin"
)

func GetImageList(c *gin.Context) {
	imageList, err := model.GetImageList()
	if err != nil {
		response.Fail(err, "Get image list fail", c)
		return
	}

	response.OK(imageList, "Get image list success", c)
}

func GetImage(c *gin.Context) {
	var image model.Image
	err := c.ShouldBindJSON(&image)
	if err != nil {
		response.Fail(err, "Invalid argument", c)
		return
	}
	
	err = image.GetImage()
	if err != nil {
		response.Fail(err, "Get image fail", c)
		return
	}
	
	response.OK(image, "Get image success", c)
}

func CreateImage(c *gin.Context) {
	var image model.Image
	err := c.ShouldBindJSON(&image)
	if err != nil {
		response.Fail(err, "Invalid argument", c)
		return
	}
	
	dockerNodeList, err := model.GetDockerNodeList()
	if err != nil {
		response.Fail(err, "Get dockerNodeList fail", c)
		return
	}

	for _, dockerNode := range dockerNodeList {
		url := "http://" + dockerNode.Host + ":" + dockerNode.Port + "/image/PullImage"
		res, err := docker.PullImage(url, image.RepoTags, image.Repository)
		if err != nil {
			response.Fail(err, "Create image fail", c)
			return
		}
		if res["code"].(float64) != 0 {
			response.Fail(err, "Create image fail", c)
			return
		}
	}

	err = image.CreateImage()
	if err != nil {
		response.Fail(err, "Create image fail", c)
		return
	}
	
	response.OK(nil, "Create image success", c)
}

func UpdateImage(c *gin.Context) {
	var image model.Image
	err := c.ShouldBindJSON(&image)
	if err != nil {
		response.Fail(err, "Invalid argument", c)
		return
	}
	
	err = image.UpdateImage()
	if err != nil {
		response.Fail(err, "Update image fail", c)
		return
	}
	
	response.OK(nil, "Update image success", c)
}

func DeleteImage(c *gin.Context) {
	var image model.Image
	err := c.ShouldBindJSON(&image)
	if err != nil {
		response.Fail(err, "Invalid argument", c)
		return
	}
	
	err = image.DeleteImage()
	if err != nil {
		response.Fail(err, "Delete image fail", c)
		return
	}
	
	response.OK(nil, "Delete image success", c)
}
