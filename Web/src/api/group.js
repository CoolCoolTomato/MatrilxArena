import {apiClient, handleResponse, handleError} from "./api.js"

function GetGroupList(data) {
  return apiClient.post("/group/GetGroupList", data).then(handleResponse).catch(handleError)
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

const groupApi = {
  GetGroupList,
  GetGroup,
  CreateGroup,
  UpdateGroup,
  DeleteGroup
}

export default groupApi
