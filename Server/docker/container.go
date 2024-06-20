package docker

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/request"
)

func GetContainerList(dockerNode model.DockerNode) (map[string]interface{}, error) {
	url := "http://" + dockerNode.Host + ":" + dockerNode.Port + PathGetContainerList
	return request.Post(url, nil)
}

func GetContainer(dockerNode model.DockerNode, containerID string) (map[string]interface{}, error) {
	url := "http://" + dockerNode.Host + ":" + dockerNode.Port + PathGetContainer
	data := map[string]interface{}{
		"ContainerID": containerID,
	}
	return request.Post(url, data)
}

func CreateContainer(dockerNode model.DockerNode, imageID string, specifiedPorts []string) (map[string]interface{}, error) {
	url := "http://" + dockerNode.Host + ":" + dockerNode.Port + PathCreateContainer
	data := map[string]interface{}{
		"ImageID":        imageID,
		"SpecifiedPorts": specifiedPorts,
	}
	return request.Post(url, data)
}

func StartContainer(dockerNode model.DockerNode, containerID string) (map[string]interface{}, error) {
	url := "http://" + dockerNode.Host + ":" + dockerNode.Port + PathStartContainer
	data := map[string]interface{}{
		"ContainerID": containerID,
	}
	return request.Post(url, data)
}

func StopContainer(dockerNode model.DockerNode, containerID string) (map[string]interface{}, error) {
	url := "http://" + dockerNode.Host + ":" + dockerNode.Port + PathStopContainer
	data := map[string]interface{}{
		"ContainerID": containerID,
	}
	return request.Post(url, data)
}

func RemoveContainer(dockerNode model.DockerNode, containerID string) (map[string]interface{}, error) {
	url := "http://" + dockerNode.Host + ":" + dockerNode.Port + PathRemoveContainer
	data := map[string]interface{}{
		"ContainerID": containerID,
	}
	return request.Post(url, data)
}

func ExecuteCommand(dockerNode model.DockerNode, containerID string, command []string) (map[string]interface{}, error) {
	url := "http://" + dockerNode.Host + ":" + dockerNode.Port + PathExecuteCommand
	data := map[string]interface{}{
		"ContainerID": containerID,
		"Command":     command,
	}
	return request.Post(url, data)
}
