import {apiClient, handleResponse, handleError} from "./api.js"

function GetUserChallengeList() {
  return apiClient.post("/userChallenge/GetUserChallengeList").then(handleResponse).catch(handleError)
}

function ResetUserChallenge(data) {
  return apiClient.post("/userChallenge/ResetUserChallenge", data).then(handleResponse).catch(handleError)
}

function CheckChallengeFlag(data) {
  return apiClient.post("/userChallenge/CheckChallengeFlag", data).then(handleResponse).catch(handleError)
}

const userChallengeApi = {
  GetUserChallengeList,
  ResetUserChallenge,
  CheckChallengeFlag
}

export default userChallengeApi
