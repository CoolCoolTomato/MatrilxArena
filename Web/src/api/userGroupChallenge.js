import {apiClient, handleResponse, handleError} from "./api.js"

function GetUserGroupChallengeList(data) {
  return apiClient.post("/userGroupChallenge/GetUserGroupChallengeList", data).then(handleResponse).catch(handleError)
}

function ResetUserGroupChallenge(data) {
  return apiClient.post("/userGroupChallenge/ResetUserGroupChallenge", data).then(handleResponse).catch(handleError)
}

function CheckGroupChallengeFlag(data) {
  return apiClient.post("/userGroupChallenge/CheckGroupChallengeFlag", data).then(handleResponse).catch(handleError)
}

const userGroupChallengeApi = {
  GetUserGroupChallengeList,
  ResetUserGroupChallenge,
  CheckGroupChallengeFlag
}

export default userGroupChallengeApi
