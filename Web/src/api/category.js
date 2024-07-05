import {apiClient, handleResponse, handleError} from "./api.js"

function GetCategoryList() {
  return apiClient.post("/category/GetCategoryList").then(handleResponse).catch(handleError)
}

function GetCategory(data) {
  return apiClient.post("/category/GetCategory", data).then(handleResponse).catch(handleError)
}

function CreateCategory(data) {
  return apiClient.post("/category/CreateCategory", data).then(handleResponse).catch(handleError)
}

function UpdateCategory(data) {
  return apiClient.post("/category/UpdateCategory", data).then(handleResponse).catch(handleError)
}

function DeleteCategory(data) {
  return apiClient.post("/category/DeleteCategory", data).then(handleResponse).catch(handleError)
}

const categoryApi = {
  GetCategoryList,
  GetCategory,
  CreateCategory,
  UpdateCategory,
  DeleteCategory
}

export default categoryApi
