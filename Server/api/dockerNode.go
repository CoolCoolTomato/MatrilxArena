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

func GetDockerNodeList(c *gin.Context) {
	dockerNodeList, err := model.GetDockerNodeList()
	if err != nil {
		response.Fail(err, "Get dockerNode list fail", c)
		return
	}
	response.OK(dockerNodeList, "Get dockerNode list success", c)
}

func GetDockerNode(c *gin.Context) {
	var dockerNode model.DockerNode
	err := c.ShouldBindJSON(&dockerNode)
	if err != nil || dockerNode.ID == 0 {
		response.Fail(err, "Invalid argument", c)
		return
	}
	
	err = dockerNode.GetDockerNode()
	if err != nil {
		response.Fail(err, "Get dockerNode fail", c)
		return
	}
	
	response.OK(dockerNode, "Get dockerNode success", c)
}

func CreateDockerNode(c *gin.Context) {
	var dockerNode model.DockerNode
	err := c.ShouldBindJSON(&dockerNode)
	if err != nil {
		response.Fail(err, "Invalid argument", c)
		return
	}
	
	err = dockerNode.CreateDockerNode()
	if err != nil {
		response.Fail(err, "Create dockerNode fail", c)
		return
	}
	
	response.OK(nil, "Create dockerNode success", c)
}

func UpdateDockerNode(c *gin.Context) {
	var dockerNode model.DockerNode
	err := c.ShouldBindJSON(&dockerNode)
	if err != nil || dockerNode.ID == 0 {
		response.Fail(err, "Invalid argument", c)
		return
	}
	
	err = dockerNode.UpdateDockerNode()
	if err != nil {
		response.Fail(err, "Update dockerNode fail", c)
		return
	}
	
	response.OK(nil, "Update dockerNode success", c)
}

func DeleteDockerNode(c *gin.Context) {
	var dockerNode model.DockerNode
	err := c.ShouldBindJSON(&dockerNode)
	if err != nil || dockerNode.ID == 0 {
		response.Fail(err, "Invalid argument", c)
		return
	}
	
	err = dockerNode.DeleteDockerNode()
	if err != nil {
		response.Fail(err, "Delete dockerNode fail", c)
		return
	}
	
	response.OK(nil, "Delete dockerNode success", c)
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

func PullImageForDockerNode(c *gin.Context) {
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

