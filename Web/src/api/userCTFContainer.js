import {apiClient, handleResponse, handleError} from "./api.js"

function GetCTFContainerListByUser() {
  return apiClient.post("/userCTFContainer/GetCTFContainerListByUser").then(handleResponse).catch(handleError)
}

function CreateCTFContainerByUser(data) {
  return apiClient.post("/userCTFContainer/CreateCTFContainerByUser", data).then(handleResponse).catch(handleError)
}

function DestroyCTFContainerByUser(data) {
  return apiClient.post("/userCTFContainer/DestroyCTFContainerByUser", data).then(handleResponse).catch(handleError)
}

function DelayCTFContainerByUser(data) {
  return apiClient.post("/userCTFContainer/DelayCTFContainerByUser", data).then(handleResponse).catch(handleError)
}

const userCTFContainerApi = {
  GetCTFContainerListByUser,
  CreateCTFContainerByUser,
  DestroyCTFContainerByUser,
  DelayCTFContainerByUser
}

export default userCTFContainerApi
