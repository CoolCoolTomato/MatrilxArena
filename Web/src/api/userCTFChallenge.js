import {apiClient, handleResponse, handleError} from "./api.js"

function GetUserCTFChallengeList(data) {
  return apiClient.post("/userCTFChallenge/GetUserCTFChallengeList", data).then(handleResponse).catch(handleError)
}

function ResetUserCTFChallenge(data) {
  return apiClient.post("/userCTFChallenge/ResetUserCTFChallenge", data).then(handleResponse).catch(handleError)
}

function CheckCTFChallengeFlag(data) {
  return apiClient.post("/userCTFChallenge/CheckCTFChallengeFlag", data).then(handleResponse).catch(handleError)
}

const userCTFChallengeApi = {
  GetUserCTFChallengeList,
  ResetUserCTFChallenge,
  CheckCTFChallengeFlag
}

export default userCTFChallengeApi
