import {apiClient, handleResponse, handleError} from "./api.js"

function GetCTFList(data) {
  return apiClient.post("/ctf/GetCTFList", data).then(handleResponse).catch(handleError)
}

function GetCTF(data) {
  return apiClient.post("/ctf/GetCTF", data).then(handleResponse).catch(handleError)
}

function CreateCTF(data) {
  return apiClient.post("/ctf/CreateCTF", data).then(handleResponse).catch(handleError)
}

function UpdateCTF(data) {
  return apiClient.post("/ctf/UpdateCTF", data).then(handleResponse).catch(handleError)
}

function DeleteCTF(data) {
  return apiClient.post("/ctf/DeleteCTF", data).then(handleResponse).catch(handleError)
}

const ctfApi = {
  GetCTFList,
  GetCTF,
  CreateCTF,
  UpdateCTF,
  DeleteCTF
}

export default ctfApi
