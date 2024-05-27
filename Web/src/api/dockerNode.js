import {apiClient, handleResponse, handleError} from "./api.js"

function GetDockerNodeList() {
    return apiClient.post("/dockerNode/GetDockerNodeList").then(handleResponse).catch(handleError)
}

function GetDockerNode(data) {
    return apiClient.post("/dockerNode/GetDockerNode", data).then(handleResponse).catch(handleError)
}

function CreateDockerNode(data) {
    return apiClient.post("/dockerNode/CreateDockerNode", data).then(handleResponse).catch(handleError)
}

function UpdateDockerNode(data) {
    return apiClient.post("/dockerNode/UpdateDockerNode", data).then(handleResponse).catch(handleError)
}

function DeleteDockerNode(data) {
    return apiClient.post("/dockerNode/DeleteDockerNode", data).then(handleResponse).catch(handleError)
}

function GetImageListFromDockerNode(data) {
    return apiClient.post("/dockerNode/GetImageListFromDockerNode", data).then(handleResponse).catch(handleError)
}

function PullImageForDockerNode(data) {
    return apiClient.post("/dockerNode/PullImageForDockerNode", data).then(handleResponse).catch(handleError)
}

function RemoveImageFromDockerNode(data) {
    return apiClient.post("/dockerNode/RemoveImageFromDockerNode", data).then(handleResponse).catch(handleError)
}
