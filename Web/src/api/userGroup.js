import {apiClient, handleResponse, handleError} from "./api.js"

function GetUserGroupList(data) {
  return apiClient.post("/userGroup/GetUserGroupList", data).then(handleResponse).catch(handleError)
}

function GetVisibleUserGroupList(data) {
  return apiClient.post("/userGroup/GetVisibleUserGroupList", data).then(handleResponse).catch(handleError)
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