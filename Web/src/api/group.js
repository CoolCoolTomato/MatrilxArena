import {apiClient, handleResponse, handleError} from "./api.js"

function GetGroupList() {
    return apiClient.post("/group/GetGroupList").then(handleResponse).catch(handleError)
}

function GetGroup(data) {
    return apiClient.post("/group/GetGroup", data).then(handleResponse).catch(handleError)
}

function CreateGroup(data) {
    return apiClient.post("/group/CreateGroup", data).then(handleResponse).catch(handleError)
}

function UpdateGroup(data) {
    return apiClient.post("/group/UpdateGroup", data).then(handleResponse).catch(handleError)
}

function DeleteGroup(data) {
    return apiClient.post("/group/DeleteGroup", data).then(handleResponse).catch(handleError)
}

function AddGroupUser(data) {
    return apiClient.post("/group/AddGroupUser", data).then(handleResponse).catch(handleError)
}

function RemoveGroupUser(data) {
    return apiClient.post("/group/RemoveGroupUser", data).then(handleResponse).catch(handleError)
}

const groupApi = {
  GetGroupList,
  GetGroup,
  CreateGroup,
  UpdateGroup,
  DeleteGroup,
  AddGroupUser,
  RemoveGroupUser
}

export default groupApi
