import {apiClient, handleResponse, handleError} from "./api.js"

function GetChallengeList(data) {
    return apiClient.post("/challenge/GetChallengeList", data).then(handleResponse).catch(handleError)
}

function GetChallenge(data) {
    return apiClient.post("/challenge/GetChallenge", data).then(handleResponse).catch(handleError)
}

function CreateChallenge(data) {
    return apiClient.post("/challenge/CreateChallenge", data).then(handleResponse).catch(handleError)
}

function UpdateChallenge(data) {
    return apiClient.post("/challenge/UpdateChallenge", data).then(handleResponse).catch(handleError)

}

function DeleteChallenge(data) {
    return apiClient.post("/challenge/DeleteChallenge", data).then(handleResponse).catch(handleError)

}

const challengeApi = {
  GetChallengeList,
  GetChallenge,
  CreateChallenge,
  UpdateChallenge,
  DeleteChallenge
}

export default challengeApi
