package api

import (
    "github.com/CoolCoolTomato/MatrilxArena/Server/common/response"
    "github.com/CoolCoolTomato/MatrilxArena/Server/docker"
    "github.com/CoolCoolTomato/MatrilxArena/Server/model"
    "github.com/gin-gonic/gin"
)

func CreateContainerFromChallenge(c *gin.Context) {
    var challenge model.Challenge
    err := c.ShouldBindJSON(&challenge)
    if err != nil || challenge.ID == 0{
        response.Fail(err, "Invalid argument", c)
        return
    }

    err = challenge.GetChallenge()
    if err != nil {
        response.Fail(err, "Get challenge fail", c)
        return
    }

    image := challenge.Image
    var dockerNode model.DockerNode
    var dockerNodeImageID string

    dockerNodeList, err := model.GetDockerNodeList()
    if err != nil {
        response.Fail(err, "Get dockerNode list fail", c)
        return
    }

    for _, dockerNode = range dockerNodeList {
        res, err := docker.GetImageList(dockerNode)
        if err != nil || res["code"].(float64) != 0 {
            continue
        }
        dockerNodeImageList := res["data"].([]interface{})
        for _, dockerNodeImageRaw := range dockerNodeImageList{
            dockerNodeImage := dockerNodeImageRaw.(map[string]interface{})
            dockerNodeImageRepoTags := dockerNodeImage["RepoTags"].([]interface{})[0]
            if dockerNodeImageRepoTags.(string) == image.RepoTags {
                dockerNodeImageID = dockerNodeImage["Id"].(string)
                break
            }
        }

    }

    res, err := docker.CreateContainer(dockerNode, dockerNodeImageID)
    if err != nil || res["code"].(float64) != 0 {
        response.Fail(err, "Create container fail", c)
        return
    }

    response.OK(res, "Create container success", c)
}
