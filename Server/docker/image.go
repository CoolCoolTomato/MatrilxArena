package docker

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/model"
	"github.com/CoolCoolTomato/MatrilxArena/Server/utils/request"
)

func GetImageList(dockerNode model.DockerNode) (map[string]interface{}, error) {
	url := "http://" + dockerNode.Host + ":" + dockerNode.Port + PathGetImageList
	return request.Post(url, nil)
}

func GetImage(dockerNode model.DockerNode, imageID string) (map[string]interface{}, error) {
	url := "http://" + dockerNode.Host + ":" + dockerNode.Port + PathGetImage
	data := map[string]interface{}{
		"ImageID": imageID,
	}
	return request.Post(url, data)
}

func PullImage(dockerNode model.DockerNode, repoTags string, repository string) (map[string]interface{}, error) {
	url := "http://" + dockerNode.Host + ":" + dockerNode.Port + PathPullImage
	data := map[string]interface{}{
		"RepoTags":   repoTags,
		"Repository": repository,
	}
	return request.Post(url, data)
}

func PullImageWithAuth(dockerNode model.DockerNode, repoTags string, repository string, username string, password string) (map[string]interface{}, error) {
	url := "http://" + dockerNode.Host + ":" + dockerNode.Port + PathPullImageWithAuth
	data := map[string]interface{}{
		"RepoTags":   repoTags,
		"Repository": repository,
		"Username":   username,
		"Password":   password,
	}
	return request.Post(url, data)
}

func RemoveImage(dockerNode model.DockerNode, imageID string) (map[string]interface{}, error) {
	url := "http://" + dockerNode.Host + ":" + dockerNode.Port + PathRemoveImage
	data := map[string]interface{}{
		"ImageID": imageID,
	}
	return request.Post(url, data)
}
