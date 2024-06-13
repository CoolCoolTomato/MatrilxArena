package docker

import (
    "github.com/CoolCoolTomato/MatrilxArena/Server/model"
    "github.com/CoolCoolTomato/MatrilxArena/Server/utils/request"
)

func GetContainerList(dockerNode model.DockerNode) (map[string]interface{}, error) {
    url := "http://" + dockerNode.Host + ":" + dockerNode.Port + PathGetContainerList
	return request.Post(url, nil)
}

func GetContainer(dockerNode model.DockerNode, ContainerID string) (map[string]interface{}, error) {
    url := "http://" + dockerNode.Host + ":" + dockerNode.Port + PathGetContainer
	data := map[string]interface{}{
		"ContainerID": ContainerID,
	}
	return request.Post(url, data)
}

func CreateContainer(dockerNode model.DockerNode, ImageID string) (map[string]interface{}, error) {
    url := "http://" + dockerNode.Host + ":" + dockerNode.Port + PathCreateContainer
	data := map[string]interface{}{
		"ImageID": ImageID,
	}
	return request.Post(url, data)
}

func StartContainer(dockerNode model.DockerNode, ContainerID string) (map[string]interface{}, error) {
    url := "http://" + dockerNode.Host + ":" + dockerNode.Port + PathStartContainer
	data := map[string]interface{}{
		"ContainerID": ContainerID,
	}
	return request.Post(url, data)
}

func StopContainer(dockerNode model.DockerNode, ContainerID string) (map[string]interface{}, error) {
    url := "http://" + dockerNode.Host + ":" + dockerNode.Port + PathStopContainer
	data := map[string]interface{}{
		"ContainerID": ContainerID,
	}
	return request.Post(url, data)
}

func RemoveContainer(dockerNode model.DockerNode, ContainerID string) (map[string]interface{}, error) {
    url := "http://" + dockerNode.Host + ":" + dockerNode.Port + PathRemoveContainer
	data := map[string]interface{}{
		"ContainerID": ContainerID,
	}
	return request.Post(url, data)
}
