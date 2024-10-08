import {apiClient, handleResponse, handleError} from "./api.js"

function GetUserList(data) {
    return apiClient.post("/user/GetUserList", data).then(handleResponse).catch(handleError)
}

function GetUser(data) {
    return apiClient.post("/user/GetUser", data).then(handleResponse).catch(handleError)
}

function CreateUser(data) {
    return apiClient.post("/user/CreateUser", data).then(handleResponse).catch(handleError)
}

function UpdateUser(data) {
    return apiClient.post("/user/UpdateUser", data).then(handleResponse).catch(handleError)
}

function DeleteUser(data) {
    return apiClient.post("/user/DeleteUser", data).then(handleResponse).catch(handleError)
}

const userApi = {
  GetUserList,
  GetUser,
  CreateUser,
  UpdateUser,
  DeleteUser
}

export default userApi
