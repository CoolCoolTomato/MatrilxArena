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
		response.Fail(err, localizer.GetMessage("DockerNodeImage.InvalidArgument", c), c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = imageRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, localizer.GetMessage("DockerNodeImage.GetDockerNodeFail", c), c)
		return
	}

	res, err := docker.GetImageList(dockerNode)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, localizer.GetMessage("DockerNodeImage.GetImageListFromDockerNodeFail", c), c)
		return
	}

	response.OK(res["data"], localizer.GetMessage("DockerNodeImage.GetImageListFromDockerNodeSuccess", c), c)
}

func GetImageFromDockerNode(c *gin.Context) {
	var imageRequest ImageRequest
	err := c.ShouldBindJSON(&imageRequest)
	if err != nil || imageRequest.DockerNodeID == 0 || imageRequest.DockerNodeImageID == "" {
		response.Fail(err, localizer.GetMessage("DockerNodeImage.InvalidArgument", c), c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = imageRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, localizer.GetMessage("DockerNodeImage.GetDockerNodeFail", c), c)
		return
	}

	res, err := docker.GetImage(dockerNode, imageRequest.DockerNodeImageID)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, localizer.GetMessage("DockerNodeImage.GetImageFromDockerNodeFail", c), c)
		return
	}

	response.OK(res["data"], localizer.GetMessage("DockerNodeImage.GetImageFromDockerNodeSuccess", c), c)
}

func PullImageFromDockerNode(c *gin.Context) {
	var imageRequest ImageRequest
	err := c.ShouldBindJSON(&imageRequest)
	if err != nil || imageRequest.DockerNodeID == 0 || imageRequest.ImageID == 0 {
		response.Fail(err, localizer.GetMessage("DockerNodeImage.InvalidArgument", c), c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = imageRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, localizer.GetMessage("DockerNodeImage.GetDockerNodeFail", c), c)
		return
	}

	image := model.Image{}
	image.ID = imageRequest.ImageID
	err = image.GetImage()
	if err != nil {
		response.Fail(err, localizer.GetMessage("DockerNodeImage.GetImageFail", c), c)
		return
	}

	res, err := docker.PullImage(dockerNode, image.RepoTags, image.Repository)
	if err != nil {
		response.Fail(err, localizer.GetMessage("DockerNodeImage.PullImageFromDockerNodeFail", c), c)
		return
	}
	if res["code"].(float64) != 0 {
		response.Fail(err, localizer.GetMessage("DockerNodeImage.PullImageFromDockerNodeFail", c), c)
		return
	}

	response.OK(image, localizer.GetMessage("DockerNodeImage.PullImageFromDockerNodeSuccess", c), c)
}

func RemoveImageFromDockerNode(c *gin.Context) {
	var imageRequest ImageRequest
	err := c.ShouldBindJSON(&imageRequest)
	if err != nil || imageRequest.DockerNodeID == 0 || imageRequest.DockerNodeImageID == "" {
		response.Fail(err, localizer.GetMessage("DockerNodeImage.InvalidArgument", c), c)
		return
	}

	var dockerNode model.DockerNode
	dockerNode.ID = imageRequest.DockerNodeID
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, localizer.GetMessage("DockerNodeImage.GetDockerNodeFail", c), c)
		return
	}

	res, err := docker.RemoveImage(dockerNode, imageRequest.DockerNodeImageID)
	if err != nil || res["code"].(float64) != 0 {
		response.Fail(err, localizer.GetMessage("DockerNodeImage.RemoveImageFromDockerNodeFail", c), c)
		return
	}

	response.OK(res["data"], localizer.GetMessage("DockerNodeImage.RemoveImageFromDockerNodeSuccess", c), c)
}
