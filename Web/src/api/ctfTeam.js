import {apiClient, handleResponse, handleError} from "./api.js"

function GetCTFTeamList(data) {
  return apiClient.post("/ctfTeam/GetCTFTeamList", data).then(handleResponse).catch(handleError)
}

function GetCTFTeam(data) {
  return apiClient.post("/ctfTeam/GetCTFTeam", data).then(handleResponse).catch(handleError)
}

function CreateCTFTeam(data) {
  return apiClient.post("/ctfTeam/CreateCTFTeam", data).then(handleResponse).catch(handleError)
}

function UpdateCTFTeam(data) {
  return apiClient.post("/ctfTeam/UpdateCTFTeam", data).then(handleResponse).catch(handleError)
}

function DeleteCTFTeam(data) {
  return apiClient.post("/ctfTeam/DeleteCTFTeam", data).then(handleResponse).catch(handleError)
}

const ctfTeamApi = {
  GetCTFTeamList,
  GetCTFTeam,
  CreateCTFTeam,
  UpdateCTFTeam,
  DeleteCTFTeam
}

export default ctfTeamApi
