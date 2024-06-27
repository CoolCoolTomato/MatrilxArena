import {apiClient, handleResponse, handleError} from "./api.js"

function GetContainerListByUser() {
  return apiClient.post("/userContainer/GetContainerListByUser").then(handleResponse).catch(handleError)
}

function CreateContainerByUser(data) {
  return apiClient.post("/userContainer/CreateContainerByUser", data).then(handleResponse).catch(handleError)
}

function DestroyContainerByUser(data) {
  return apiClient.post("/userContainer/DestroyContainerByUser", data).then(handleResponse).catch(handleError)
}

function DelayContainerByUser(data) {
  return apiClient.post("/userContainer/DelayContainerByUser", data).then(handleResponse).catch(handleError)
}

const userContainerApi = {
  GetContainerListByUser,
  CreateContainerByUser,
  DestroyContainerByUser,
  DelayContainerByUser
}

export default userContainerApi
