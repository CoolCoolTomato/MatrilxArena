import {apiClient, handleResponse, handleError} from "./api.js"

function GetUserGroupList() {
  return apiClient.post("/userGroup/GetUserGroupList").then(handleResponse).catch(handleError)
}

function GetVisibleUserGroupList() {
  return apiClient.post("/userGroup/GetVisibleUserGroupList").then(handleResponse).catch(handleError)
}

function AddUserGroup(data) {
  return apiClient.post("/userGroup/AddUserGroup", data).then(handleResponse).catch(handleError)
}

function RemoveUserGroup(data) {
  return apiClient.post("/userGroup/RemoveUserGroup", data).then(handleResponse).catch(handleError)
}

const userGroupApi = {
  GetUserGroupList,
  GetVisibleUserGroupList,
  AddUserGroup,
  RemoveUserGroup
}

export default userGroupApi