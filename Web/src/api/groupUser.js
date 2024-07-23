import {apiClient, handleResponse, handleError} from "./api.js"

function GetGroupUserList(data) {
  return apiClient.post("/groupUser/GetGroupUserList", data).then(handleResponse).catch(handleError)
}

function AddGroupUser(data) {
    return apiClient.post("/groupUser/AddGroupUser", data).then(handleResponse).catch(handleError)
}

function RemoveGroupUser(data) {
    return apiClient.post("/groupUser/RemoveGroupUser", data).then(handleResponse).catch(handleError)
}

const groupUserApi = {
  GetGroupUserList,
  AddGroupUser,
  RemoveGroupUser
}

export default groupUserApi
