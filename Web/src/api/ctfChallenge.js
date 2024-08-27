import {apiClient, handleResponse, handleError} from "./api.js"

function GetCTFChallengeList(data) {
  return apiClient.post("/ctfChallenge/GetCTFChallengeList", data).then(handleResponse).catch(handleError)
}

function GetCTFChallenge(data) {
  return apiClient.post("/ctfChallenge/GetCTFChallenge", data).then(handleResponse).catch(handleError)
}

function CreateCTFChallenge(data) {
  return apiClient.post("/ctfChallenge/CreateCTFChallenge", data).then(handleResponse).catch(handleError)
}

function UpdateCTFChallenge(data) {
  return apiClient.post("/ctfChallenge/UpdateCTFChallenge", data).then(handleResponse).catch(handleError)
}

function DeleteCTFChallenge(data) {
  return apiClient.post("/ctfChallenge/DeleteCTFChallenge", data).then(handleResponse).catch(handleError)
}

const ctfChallengeApi = {
  GetCTFChallengeList,
  GetCTFChallenge,
  CreateCTFChallenge,
  UpdateCTFChallenge,
  DeleteCTFChallenge
}

export default ctfChallengeApi
