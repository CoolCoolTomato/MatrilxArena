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

export function DeleteDockerNode(data) {
    return apiClient.post("/dockerNode/DeleteDockerNode", data).then(handleResponse).catch(handleError)
}

function GetImageListFromDockerNode(data) {
    return apiClient.post("/dockerNode/GetImageListFromDockerNode", data).then(handleResponse).catch(handleError)
}

function GetImageFromDockerNode(data) {
    return apiClient.post("/dockerNode/GetImageFromDockerNode", data).then(handleResponse).catch(handleError)
}

function PullImageFromDockerNode(data) {
    return apiClient.post("/dockerNode/PullImageFromDockerNode", data).then(handleResponse).catch(handleError)
}

function RemoveImageFromDockerNode(data) {
    return apiClient.post("/dockerNode/RemoveImageFromDockerNode", data).then(handleResponse).catch(handleError)
}

function GetContainerListFromDockerNode(data) {
    return apiClient.post("/dockerNode/GetContainerListFromDockerNode", data).then(handleResponse).catch(handleError)
}

function GetContainerFromDockerNode(data) {
    return apiClient.post("/dockerNode/GetContainerFromDockerNode", data).then(handleResponse).catch(handleError)
}

function CreateContainerFromDockerNode(data) {
    return apiClient.post("/dockerNode/CreateContainerFromDockerNode", data).then(handleResponse).catch(handleError)
}

function StartContainerFromDockerNode(data) {
    return apiClient.post("/dockerNode/StartContainerFromDockerNode", data).then(handleResponse).catch(handleError)
}

function StopContainerFromDockerNode(data) {
    return apiClient.post("/dockerNode/StopContainerFromDockerNode", data).then(handleResponse).catch(handleError)
}

function RemoveContainerFromDockerNode(data) {
    return apiClient.post("/dockerNode/RemoveContainerFromDockerNode", data).then(handleResponse).catch(handleError)
}

const dockerNodeApi = {
    GetDockerNodeList,
    GetDockerNode,
    CreateDockerNode,
    UpdateDockerNode,
    DeleteDockerNode,

    GetImageListFromDockerNode,
    GetImageFromDockerNode,
    PullImageFromDockerNode,
    RemoveImageFromDockerNode,

    GetContainerListFromDockerNode,
    GetContainerFromDockerNode,
    CreateContainerFromDockerNode,
    StartContainerFromDockerNode,
    StopContainerFromDockerNode,
    RemoveContainerFromDockerNode
}
export default dockerNodeApi
