package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/DockerNode/common/response"
	"github.com/CoolCoolTomato/MatrilxArena/DockerNode/docker"
	"github.com/gin-gonic/gin"
)

type ImageRequest struct {
	ImageID string
	ImageName string
	Username string
	Password string
}

func GetImageList(c *gin.Context) {
	images, err := docker.GetImageList()
	if err != nil {
		response.Fail(err, "Get image list fail", c)
		return
	}
	
	response.OK(images, "Get image list success", c)
}

func GetImage(c *gin.Context) {
	var imageRequest ImageRequest
	if err := c.ShouldBindJSON(&imageRequest); err != nil {
		response.Fail(err, "Invalid argument", c)
		return
	}
	
	image, err := docker.GetImage(imageRequest.ImageID)
	if err != nil {
		response.Fail(err, "Get image fail", c)
		return
	}
	
	response.OK(image, "Get iamge success", c)
}

func PullImage(c *gin.Context) {
	var imageRequest ImageRequest
	if err := c.ShouldBindJSON(&imageRequest); err != nil {
		response.Fail(err, "Invalid argument", c)
		return
	}
	
	err := docker.PullImage(imageRequest.ImageName)
	if err != nil {
		response.Fail(err, "Pull image fail", c)
		return
	}
	
	response.OK(nil, "Pull iamge success", c)
}

func PullImageWithAuth(c *gin.Context) {
	var imageRequest ImageRequest
	if err := c.ShouldBindJSON(&imageRequest); err != nil {
		response.Fail(err, "Invalid argument", c)
		return
	}
	
	err := docker.PullImageWithAuth(imageRequest.ImageName, imageRequest.Username, imageRequest.Password)
	if err != nil {
		response.Fail(err, "Pull image fail", c)
		return
	}
	
	response.OK(nil, "Pull iamge success", c)
}

func RemoveImage(c *gin.Context) {
	var imageRequest ImageRequest
	if err := c.ShouldBindJSON(&imageRequest); err != nil {
		response.Fail(err, "Invalid argument", c)
		return
	}
	
	err := docker.RemoveImage(imageRequest.ImageID)
	if err != nil {
		response.Fail(err, "Remove image fail", c)
		return
	}
	
	response.OK(nil, "Remove iamge success", c)
}
