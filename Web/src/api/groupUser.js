import {apiClient, handleResponse, handleError} from "./api.js"

function AddGroupUser(data) {
    return apiClient.post("/groupUser/AddGroupUser", data).then(handleResponse).catch(handleError)
}

function RemoveGroupUser(data) {
    return apiClient.post("/groupUser/RemoveGroupUser", data).then(handleResponse).catch(handleError)
}

const groupUserApi = {
  AddGroupUser,
  RemoveGroupUser
}

export default groupUserApi
