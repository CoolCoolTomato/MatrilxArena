import {apiClient, handleResponse, handleError} from "./api.js"

function GetGroupChallengeList(data) {
    return apiClient.post("/groupChallenge/GetGroupChallengeList", data).then(handleResponse).catch(handleError)
}

function GetGroupChallenge(data) {
    return apiClient.post("/groupChallenge/GetGroupChallenge", data).then(handleResponse).catch(handleError)
}

function CreateGroupChallenge(data) {
    return apiClient.post("/groupChallenge/CreateGroupChallenge", data).then(handleResponse).catch(handleError)
}

function UpdateGroupChallenge(data) {
    return apiClient.post("/groupChallenge/UpdateGroupChallenge", data).then(handleResponse).catch(handleError)

}

function DeleteGroupChallenge(data) {
    return apiClient.post("/groupChallenge/DeleteGroupChallenge", data).then(handleResponse).catch(handleError)

}

const groupChallengeApi = {
  GetGroupChallengeList,
  GetGroupChallenge,
  CreateGroupChallenge,
  UpdateGroupChallenge,
  DeleteGroupChallenge
}

export default groupChallengeApi
