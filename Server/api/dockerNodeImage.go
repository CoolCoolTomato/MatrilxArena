package api

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/docker"
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/localizer"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
	"github.com/gin-gonic/gin"
)

type ImageRequest struct {
	DockerNodeID      uint
	DockerNodeImageID string
	ImageID           uint
}

func GetImageListFromDockerNode(c *gin.Context) {
	var imageRequest ImageRequest
	err := c.ShouldBindJSON(&imageRequest)
	if err != nil || imageRequest.DockerNodeID == 0 {
		response.Fail(err, localizer.GetMessage("InvalidArgument", c), c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = imageRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, localizer.GetMessage("GetDockerNodeFail", c), c)
		return
	}

	res, err := docker.GetImageList(dockerNode)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, localizer.GetMessage("GetImageListFromDockerNodeFail", c), c)
		return
	}

	response.OK(res["data"], localizer.GetMessage("GetImageListFromDockerNodeSuccess", c), c)
}

func GetImageFromDockerNode(c *gin.Context) {
	var imageRequest ImageRequest
	err := c.ShouldBindJSON(&imageRequest)
	if err != nil || imageRequest.DockerNodeID == 0 || imageRequest.DockerNodeImageID == "" {
		response.Fail(err, localizer.GetMessage("InvalidArgument", c), c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = imageRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, localizer.GetMessage("GetDockerNodeFail", c), c)
		return
	}

	res, err := docker.GetImage(dockerNode, imageRequest.DockerNodeImageID)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, localizer.GetMessage("GetImageFromDockerNodeFail", c), c)
		return
	}

	response.OK(res["data"], localizer.GetMessage("GetImageFromDockerNodeSuccess", c), c)
}

func PullImageFromDockerNode(c *gin.Context) {
	var imageRequest ImageRequest
	err := c.ShouldBindJSON(&imageRequest)
	if err != nil || imageRequest.DockerNodeID == 0 || imageRequest.ImageID == 0 {
		response.Fail(err, localizer.GetMessage("InvalidArgument", c), c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = imageRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, localizer.GetMessage("GetDockerNodeFail", c), c)
		return
	}

	image := model.Image{}
	image.ID = imageRequest.ImageID
	err = image.GetImage()
	if err != nil {
		response.Fail(err, localizer.GetMessage("GetImageFail", c), c)
		return
	}

	res, err := docker.PullImage(dockerNode, image.RepoTags, image.Repository)
	if err != nil {
		response.Fail(err, localizer.GetMessage("PullImageFromDockerNodeFail", c), c)
		return
	}
	if res["code"].(float64) != 0 {
		response.Fail(err, localizer.GetMessage("PullImageFromDockerNodeFail", c), c)
		return
	}

	response.OK(image, localizer.GetMessage("PullImageFromDockerNodeSuccess", c), c)
}

func RemoveImageFromDockerNode(c *gin.Context) {
	var imageRequest ImageRequest
	err := c.ShouldBindJSON(&imageRequest)
	if err != nil || imageRequest.DockerNodeID == 0 || imageRequest.DockerNodeImageID == "" {
		response.Fail(err, localizer.GetMessage("InvalidArgument", c), c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = imageRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, localizer.GetMessage("GetDockerNodeFail", c), c)
		return
	}

	res, err := docker.RemoveImage(dockerNode, imageRequest.DockerNodeImageID)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, localizer.GetMessage("RemoveImageFromDockerNodeFail", c), c)
		return
	}

	response.OK(res["data"], localizer.GetMessage("RemoveImageFromDockerNodeSuccess", c), c)
}
