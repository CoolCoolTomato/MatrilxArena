import {apiClient, handleResponse, handleError} from "./api.js"

function GetAttachmentList() {
  return apiClient.post("/attachment/GetAttachmentList").then(handleResponse).catch(handleError)
}

function GetAttachment(data) {
  return apiClient.post("/attachment/GetAttachment", data).then(handleResponse).catch(handleError)
}

function CreateAttachment(data) {
  return apiClient.post("/attachment/CreateAttachment", data).then(handleResponse).catch(handleError)
}

function UpdateAttachment(data) {
  return apiClient.post("/attachment/UpdateAttachment", data).then(handleResponse).catch(handleError)
}

function DeleteAttachment(data) {
  return apiClient.post("/attachment/DeleteAttachment", data).then(handleResponse).catch(handleError)
}

function UploadAttachment(fileName, file) {
  const formData = new FormData();
  formData.append('file', file);
  if (fileName) {
    formData.append('fileName', fileName);
  }

  return apiClient.post("/attachment/UploadAttachment", formData, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  }).then(handleResponse).catch(handleError);
}

const attachmentApi = {
  GetAttachmentList, GetAttachment,
  CreateAttachment,
  UpdateAttachment,
  DeleteAttachment,
  UploadAttachment
}

export default attachmentApi
