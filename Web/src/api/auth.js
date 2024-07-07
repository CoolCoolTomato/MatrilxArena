import {apiClient, handleResponse, handleError} from "./api.js"

function Login(data) {
  return apiClient.post("/auth/login", data).then(handleResponse).catch(handleError)
}

function GetUserByAuth() {
  return apiClient.post("/auth/GetUserByAuth").then(handleResponse).catch(handleError)
}

const authApi = {
  Login,
  GetUserByAuth
}

export default authApi
