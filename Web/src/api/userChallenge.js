import {apiClient, handleResponse, handleError} from "./api.js"

function GetUserContainer(data) {
  return apiClient.post("/userChallenge/GetUserContainer", data).then(handleResponse).catch(handleError)
}

function CheckFlag(data) {
  return apiClient.post("/userChallenge/CheckFlag", data).then(handleResponse).catch(handleError)
}

export default {CheckFlag}