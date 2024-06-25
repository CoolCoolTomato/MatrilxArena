import {apiClient, handleResponse, handleError} from "./api.js"

function CheckFlag(data) {
  return apiClient.post("/userChallenge/CheckFlag", data).then(handleResponse).catch(handleError)
}

export default {CheckFlag}