import {apiClient, handleResponse, handleError} from "./api.js"

function GetGroupContainerListByUser() {
  return apiClient.post("/userGroupContainer/GetGroupContainerListByUser").then(handleResponse).catch(handleError)
}

function CreateGroupContainerByUser(data) {
  return apiClient.post("/userGroupContainer/CreateGroupContainerByUser", data).then(handleResponse).catch(handleError)
}

function DestroyGroupContainerByUser(data) {
  return apiClient.post("/userGroupContainer/DestroyGroupContainerByUser", data).then(handleResponse).catch(handleError)
}

function DelayGroupContainerByUser(data) {
  return apiClient.post("/userGroupContainer/DelayGroupContainerByUser", data).then(handleResponse).catch(handleError)
}

const userGroupContainerApi = {
  GetGroupContainerListByUser,
  CreateGroupContainerByUser,
  DestroyGroupContainerByUser,
  DelayGroupContainerByUser
}

export default userGroupContainerApi
