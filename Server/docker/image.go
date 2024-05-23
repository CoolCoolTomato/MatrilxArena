package docker

import "github.com/CoolCoolTomato/MatrilxArena/Server/common/request"

func GetImageList(url string) (map[string]interface{}, error) {
	return request.Post(url, nil)
}

func GetImage(url string, ImageID string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"ImageID": ImageID,
	}
	return request.Post(url, data)
}

func PullImage(url string, ImageName string, Repository string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"ImageName": ImageName,
		"Repository": Repository,
	}
	return request.Post(url, data)
}

func PullImageWithAuth(url string, ImageName string, Repository string, Username string, Password string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"ImageName": ImageName,
		"Repository": Repository,
		"Username": Username,
		"Password": Password,
	}
	return request.Post(url, data)
}

func RemoveImage(url string, ImageID string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"ImageID": ImageID,
	}
	return request.Post(url, data)
}