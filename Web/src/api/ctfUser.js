import {apiClient, handleResponse, handleError} from "./api.js"

function GetCTFUserList(data) {
  return apiClient.post("/ctfUser/GetCTFUserList", data).then(handleResponse).catch(handleError)
}

function AddCTFUser(data) {
    return apiClient.post("/ctfUser/AddCTFUser", data).then(handleResponse).catch(handleError)
}

function RemoveCTFUser(data) {
    return apiClient.post("/ctfUser/RemoveCTFUser", data).then(handleResponse).catch(handleError)
}

const ctfUserApi = {
  GetCTFUserList,
  AddCTFUser,
  RemoveCTFUser
}

export default ctfUserApi
