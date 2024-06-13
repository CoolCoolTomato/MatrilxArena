package api

import (
    "github.com/CoolCoolTomato/MatrilxArena/Server/docker"
    "github.com/CoolCoolTomato/MatrilxArena/Server/model"
    "github.com/CoolCoolTomato/MatrilxArena/Server/utils/manager"
    "github.com/CoolCoolTomato/MatrilxArena/Server/utils/response"
    "github.com/gin-gonic/gin"
)

func CreateContainerFromChallenge(c *gin.Context) {
    username, exists := c.Get("Username")
    if !exists {
        response.Fail(nil, "User not found", c)
        return
    }

    userContainers, err := manager.GetUserContainers(username.(string))
    if err != nil {
        response.Fail(nil, "Get user containers fail", c)
        return
    }
    if len(userContainers) >= 3 {
        response.Fail(nil, "You can only create up to 3 containers", c)
        return
    }

    var challenge model.Challenge
    err = c.ShouldBindJSON(&challenge)
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

    containerID := res["data"].(string)
    err = manager.AddUserContainer(username.(string), containerID, dockerNode.ID)

    container, err := docker.GetContainer(dockerNode, containerID)
    if err != nil || res["code"].(float64) != 0 {
        response.Fail(err, "Get container fail", c)
        return
    }
    response.OK(container["data"], "Create container success", c)
}
