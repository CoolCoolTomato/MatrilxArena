import {apiClient, handleResponse, handleError} from "./api.js"

function Login(data) {
  return apiClient.post("/auth/login", data).then(handleResponse).catch(handleError)
}

const authApi = {
  Login
}

export default authApi