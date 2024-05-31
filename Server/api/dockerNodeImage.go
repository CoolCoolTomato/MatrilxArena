package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/common/response"
	"github.com/CoolCoolTomato/MatrilxArena/Server/docker"
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/gin-gonic/gin"
)

type ImageRequest struct {
	DockerNodeID uint
	DockerNodeImageID string
    ImageID uint
}

func GetImageListFromDockerNode(c *gin.Context) {
	var imageRequest ImageRequest
	err := c.ShouldBindJSON(&imageRequest)
	if err != nil || imageRequest.DockerNodeID == 0 {
		response.Fail(err, "Invalid argument", c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = imageRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, "Get dockerNode fail", c)
		return
	}

	url := "http://" + dockerNode.Host + ":" + dockerNode.Port + docker.PathGetImageList
	res, err := docker.GetImageList(url)
	if err != nil || res["code"].(float64) != 0{
		response.Fail(err, "Get image list fail", c)
		return
	}

	response.OK(res["data"], "Get image list success", c)
}

func GetImageFromDockerNode(c *gin.Context) {
	var imageRequest ImageRequest
	err := c.ShouldBindJSON(&imageRequest)
	if err != nil || imageRequest.DockerNodeID == 0 || imageRequest.DockerNodeImageID == "" {
		response.Fail(err, "Invalid argument", c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = imageRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, "Get dockerNode fail", c)
		return
	}

	url := "http://" + dockerNode.Host + ":" + dockerNode.Port + docker.PathGetImage
	res, err := docker.GetImage(url, imageRequest.DockerNodeImageID)
	if err != nil || res["code"].(float64) != 0{
		response.Fail(err, "Get image fail", c)
		return
	}

	response.OK(res["data"], "Get image success", c)
}

func PullImageFromDockerNode(c *gin.Context) {
	var imageRequest ImageRequest
	err := c.ShouldBindJSON(&imageRequest)
	if err != nil || imageRequest.DockerNodeID == 0 || imageRequest.ImageID == 0 {
		response.Fail(err, "Invalid argument", c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = imageRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, "Get dockerNode fail", c)
		return
	}

	image := model.Image{}
	image.ID = imageRequest.ImageID
	err = image.GetImage()
    if err != nil {
        response.Fail(err, "Get image fail", c)
        return
    }

    url := "http://" + dockerNode.Host + ":" + dockerNode.Port + docker.PathPullImage
    res, err := docker.PullImage(url, image.RepoTags, image.Repository)
    if err != nil {
        response.Fail(err, "Pull image fail", c)
        return
    }
    if res["code"].(float64) != 0 {
        response.Fail(err, "Pull image fail", c)
        return
    }

    response.OK(image, "Pull image success", c)
}

func RemoveImageFromDockerNode(c *gin.Context) {
	var imageRequest ImageRequest
	err := c.ShouldBindJSON(&imageRequest)
	if err != nil || imageRequest.DockerNodeID == 0 || imageRequest.DockerNodeImageID == ""{
		response.Fail(err, "Invalid argument", c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = imageRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, "Get dockerNode fail", c)
		return
	}

	url := "http://" + dockerNode.Host + ":" + dockerNode.Port + docker.PathRemoveImage
	res, err := docker.RemoveImage(url, imageRequest.DockerNodeImageID)
	if err != nil || res["code"].(float64) != 0{
		response.Fail(err, "Remove image fail", c)
		return
	}

	response.OK(res["data"], "Remove image success", c)
}
