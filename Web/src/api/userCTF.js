import {apiClient, handleResponse, handleError} from "./api.js"

function GetUserCTFList(data) {
  return apiClient.post("/userCTF/GetUserCTFList", data).then(handleResponse).catch(handleError)
}

function GetVisibleUserCTFList(data) {
  return apiClient.post("/userCTF/GetVisibleUserCTFList", data).then(handleResponse).catch(handleError)
}

function AddUserCTF(data) {
  return apiClient.post("/userCTF/AddUserCTF", data).then(handleResponse).catch(handleError)
}

function RemoveUserCTF(data) {
  return apiClient.post("/userCTF/RemoveUserCTF", data).then(handleResponse).catch(handleError)
}

const userCTFApi = {
  GetUserCTFList,
  GetVisibleUserCTFList,
  AddUserCTF,
  RemoveUserCTF
}

export default userCTFApi
