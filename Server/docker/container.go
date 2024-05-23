package docker

import "github.com/CoolCoolTomato/MatrilxArena/Server/common/request"

func GetContainerList(url string) (map[string]interface{}, error) {
	return request.Post(url, nil)
}

func GetContainer(url string, ContainerID string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"ContainerID": ContainerID,
	}
	return request.Post(url, data)
}

func CreateContainer(url string, ImageID string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"ImageID": ImageID,
	}
	return request.Post(url, data)
}

func StartContainer(url string, ContainerID string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"ContainerID": ContainerID,
	}
	return request.Post(url, data)
}

func StopContainer(url string, ContainerID string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"ContainerID": ContainerID,
	}
	return request.Post(url, data)
}

func RemoveContainer(url string, ContainerID string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"ContainerID": ContainerID,
	}
	return request.Post(url, data)
}