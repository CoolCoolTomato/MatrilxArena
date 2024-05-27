import {apiClient, handleResponse, handleError} from "./api.js"

function GetImageList() {
    return apiClient.post("/image/GetImageList").then(handleResponse).catch(handleError)
}

function GetImage(data) {
    return apiClient.post("/image/GetImage", data).then(handleResponse).catch(handleError)
}

function CreateImage(data) {
    return apiClient.post("/image/CreateImage", data).then(handleResponse).catch(handleError)
}

function UpdateImage(data) {
    return apiClient.post("/image/UpdateImage", data).then(handleResponse).catch(handleError)
}

function DeleteImage(data) {
    return apiClient.post("/image/DeleteImage", data).then(handleResponse).catch(handleError)
}
