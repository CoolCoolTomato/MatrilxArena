package docker

const (
	PathGetImageList = "/image/GetImageList"
	PathGetImage = "/image/GetImage"
	PathPullImage = "/image/PullImage"
	PathPullImageWithAuth = "/image/PullImageWithAuth"
	PathRemoveImage = "/image/RemoveImage"
	
	PathGetContainerList = "/container/GetContainerList"
	PathGetContainer = "/container/GetContainer"
	PathCreateContainer = "/container/CreateContainer"
	PathStartContainer = "/container/StartContainer"
	PathStopContainer = "/container/StopContainer"
	PathRemoveContainer = "/container/RemoveContainer"
)
