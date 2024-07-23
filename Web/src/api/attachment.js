import {apiClient, handleResponse, handleError} from "./api.js"

function GetAttachmentList(data) {
  return apiClient.post("/attachment/GetAttachmentList", data).then(handleResponse).catch(handleError)
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
  const formData = new FormData()
  formData.append('file', file)
  if (fileName) {
    formData.append('fileName', fileName)
  }

  return apiClient.post("/attachment/UploadAttachment", formData, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  }).then(handleResponse).catch(handleError)
}

function DownloadAttachment(id) {
  return apiClient.get(`/attachment/DownloadAttachment/${id}`, {
    responseType: 'blob',
  }).then(response => {
    if (response.data.type === "application/octet-stream") {
      const contentDisposition = response.headers['content-disposition']
      let fileName = `attachment${id}`
      if (contentDisposition) {
        const fileNameMatch = contentDisposition.match(/filename="?(.+)"?/i)
        if (fileNameMatch && fileNameMatch[1]) {
          fileName = fileNameMatch[1].replace(/['"]/g, '')
        }
      }
      const contentType = response.headers['content-type'] || 'application/octet-stream'
      const blob = new Blob([response.data], { type: contentType })
      const url = window.URL.createObjectURL(blob)
      const link = document.createElement('a')
      link.href = url
      link.download = fileName
      document.body.appendChild(link)
      link.click()
      document.body.removeChild(link)
      window.URL.revokeObjectURL(url)
      return true
    } else {
      return false
    }
  }).catch(handleError)
}

const attachmentApi = {
  GetAttachmentList,
  GetAttachment,
  CreateAttachment,
  UpdateAttachment,
  DeleteAttachment,
  UploadAttachment,
  DownloadAttachment,
}

export default attachmentApi
