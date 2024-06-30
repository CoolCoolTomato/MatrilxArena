import {apiClient, handleResponse, handleError} from "./api.js"

function Login(data) {
  return apiClient.post("/auth/login", data).then(handleResponse).catch(handleError)
}

function CheckAuth() {
  return apiClient.post("/auth/CheckAuth").then(handleResponse).catch(handleError)
}

const authApi = {
  Login,
  CheckAuth
}

export default authApi
