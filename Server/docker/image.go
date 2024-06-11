package docker

import (
    "github.com/CoolCoolTomato/MatrilxArena/Server/common/request"
    "github.com/CoolCoolTomato/MatrilxArena/Server/model"
)

func GetImageList(dockerNode model.DockerNode) (map[string]interface{}, error) {
    url := "http://" + dockerNode.Host + ":" + dockerNode.Port + PathGetImageList
	return request.Post(url, nil)
}

func GetImage(dockerNode model.DockerNode, ImageID string) (map[string]interface{}, error) {
    url := "http://" + dockerNode.Host + ":" + dockerNode.Port + PathGetImage
	data := map[string]interface{}{
		"ImageID": ImageID,
	}
	return request.Post(url, data)
}

func PullImage(dockerNode model.DockerNode, RepoTags string, Repository string) (map[string]interface{}, error) {
    url := "http://" + dockerNode.Host + ":" + dockerNode.Port + PathPullImage
	data := map[string]interface{}{
		"RepoTags": RepoTags,
		"Repository": Repository,
	}
	return request.Post(url, data)
}

func PullImageWithAuth(dockerNode model.DockerNode, RepoTags string, Repository string, Username string, Password string) (map[string]interface{}, error) {
    url := "http://" + dockerNode.Host + ":" + dockerNode.Port + PathPullImageWithAuth
    data := map[string]interface{}{
		"RepoTags": RepoTags,
		"Repository": Repository,
		"Username": Username,
		"Password": Password,
	}
	return request.Post(url, data)
}

func RemoveImage(dockerNode model.DockerNode, ImageID string) (map[string]interface{}, error) {
    url := "http://" + dockerNode.Host + ":" + dockerNode.Port + PathRemoveImage
	data := map[string]interface{}{
		"ImageID": ImageID,
	}
	return request.Post(url, data)
}
