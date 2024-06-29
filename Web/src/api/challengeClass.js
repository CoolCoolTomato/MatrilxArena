import {apiClient, handleResponse, handleError} from "./api.js"

function GetChallengeClassList() {
  return apiClient.post("/challengeClass/GetChallengeClassList").then(handleResponse).catch(handleError)
}

function GetChallengeClass(data) {
  return apiClient.post("/challengeClass/GetChallengeClass", data).then(handleResponse).catch(handleError)
}

function CreateChallengeClass(data) {
  return apiClient.post("/challengeClass/CreateChallengeClass", data).then(handleResponse).catch(handleError)
}

function UpdateChallengeClass(data) {
  return apiClient.post("/challengeClass/UpdateChallengeClass", data).then(handleResponse).catch(handleError)
}

function DeleteChallengeClass(data) {
  return apiClient.post("/challengeClass/DeleteChallengeClass", data).then(handleResponse).catch(handleError)
}

const challengeClassApi = {
  GetChallengeClassList,
  GetChallengeClass,
  CreateChallengeClass,
  UpdateChallengeClass,
  DeleteChallengeClass
}

export default challengeClassApi
