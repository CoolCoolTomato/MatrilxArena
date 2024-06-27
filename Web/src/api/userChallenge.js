import {apiClient, handleResponse, handleError} from "./api.js"

function GetUserChallengeList() {
  return apiClient.post("/userChallenge/GetUserChallengeList").then(handleResponse).catch(handleError)
}

function ResetUserChallenge(data) {
  return apiClient.post("/userChallenge/ResetUserChallenge", data).then(handleResponse).catch(handleError)
}

function CheckFlag(data) {
  return apiClient.post("/userChallenge/CheckFlag", data).then(handleResponse).catch(handleError)
}

const userChallengeApi = {
  GetUserChallengeList,
  ResetUserChallenge,
  CheckFlag
}

export default userChallengeApi