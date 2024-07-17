import {apiClient, handleResponse, handleError} from "./api.js"

function GetGroupChallengeList() {
    return apiClient.post("/groupChallenge/GetGroupChallengeList").then(handleResponse).catch(handleError)
}

function GetGroupChallengeListByQuery(data) {
    return apiClient.post("/groupChallenge/GetGroupChallengeListByQuery", data).then(handleResponse).catch(handleError)
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
  GetGroupChallengeListByQuery,
  GetGroupChallenge,
  CreateGroupChallenge,
  UpdateGroupChallenge,
  DeleteGroupChallenge
}

export default groupChallengeApi
