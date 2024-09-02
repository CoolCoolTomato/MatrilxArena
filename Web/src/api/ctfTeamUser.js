import {apiClient, handleResponse, handleError} from "./api.js"

function GetCTFTeamUserList(data) {
  return apiClient.post("/ctfTeamUser/GetCTFTeamUserList", data).then(handleResponse).catch(handleError)
}

function AddCTFTeamUser(data) {
    return apiClient.post("/ctfTeamUser/AddCTFTeamUser", data).then(handleResponse).catch(handleError)
}

function RemoveCTFTeamUser(data) {
    return apiClient.post("/ctfTeamUser/RemoveCTFTeamUser", data).then(handleResponse).catch(handleError)
}

const ctfTeamUserApi = {
  GetCTFTeamUserList,
  AddCTFTeamUser,
  RemoveCTFTeamUser
}

export default ctfTeamUserApi
