import {apiClient, handleResponse, handleError} from "./api.js"

function GetChallengeList() {
    return apiClient.post("/challenge/GetChallengeList").then(handleResponse).catch(handleError)
}

function GetChallenge(data) {
    return apiClient.post("/challenge/GetChallenge", data).then(handleResponse).catch(handleError)
}

function CreateChallenge(data) {
    return apiClient.post("/challenge/GetChallenge", data).then(handleResponse).catch(handleError)
}

function UpdateChallenge(data) {
    return apiClient.post("/challenge/UpdateChallenge", data).then(handleResponse).catch(handleError)

}

function DeleteChallenge(data) {
    return apiClient.post("/challenge/DeleteChallenge", data).then(handleResponse).catch(handleError)

}

export default {GetChallengeList, GetChallenge, CreateChallenge, UpdateChallenge, DeleteChallenge}