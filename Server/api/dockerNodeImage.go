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
	ImageIDList []uint
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
	if err != nil || imageRequest.DockerNodeID == 0 || len(imageRequest.ImageIDList) == 0 {
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

	var imageList []model.Image
	for _, id := range imageRequest.ImageIDList {
		image := model.Image{}
		image.ID = id
		err = image.GetImage()
		if err != nil {
			continue
		}
		imageList = append(imageList, image)
	}

	for _, image := range imageList {
		url := "http://" + dockerNode.Host + ":" + dockerNode.Port + docker.PathPullImage
		res, err := docker.PullImage(url, image.RepoTags, image.Repository)
		if err != nil {
			response.Fail(err, "Pull image fail", c)
			continue
		}
		if res["code"].(float64) != 0 {
			response.Fail(err, "Pull image fail", c)
			continue
		}
	}

	response.OK(nil, "Pull image success", c)
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