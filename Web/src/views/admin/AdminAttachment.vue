<template>
  <el-container>
    <el-header>
      <h2 style="color: var(--el-text-color-primary)">Attachment Manager</h2>
    </el-header>
    <el-main>
      <el-button style="margin: 10px" @click="createAttachmentFormVisible = true">Add</el-button>
      <el-button style="margin: 10px" @click="uploadAttachmentFormVisible = true">Upload</el-button>
      <el-table
        :data="attachmentList"
        table-layout="fixed"
      >
        <el-table-column prop="FileName" label="FileName"/>
        <el-table-column prop="FilePath" label="FilePath"/>
        <el-table-column fixed="right" label="Operations" width="320px">
          <template #default=scope>
            <el-button
              @click="OpenUpdateAttachmentForm(scope.row)"
            >
              Update
            </el-button>
            <el-button
              @click="OpenDeleteAttachmentForm(scope.row)"
            >
              Delete
            </el-button>
            <el-button
              @click="DownloadAttachment(scope.row.ID)"
              >
              Download
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-dialog
        v-model="createAttachmentFormVisible"
        title="Create Attachment"
        width="500"
        @close="ClearCreateAttachmentForm"
      >
        <el-form :model=createAttachmentData>
          <el-form-item label="FileName" :label-width="labelWidth">
            <el-input v-model="createAttachmentData.FileName"/>
          </el-form-item>
          <el-form-item label="FilePath" :label-width="labelWidth">
            <el-input v-model="createAttachmentData.FilePath"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="createAttachmentFormVisible = false">Cancel</el-button>
          <el-button @click="CreateAttachment">Submit</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="updateAttachmentFormVisible"
        title="Update Attachment"
        width="500"
        @close="ClearUpdateAttachmentForm"
      >
        <el-form :model=updateAttachmentData>
          <el-form-item label="FileName" :label-width="labelWidth">
            <el-input v-model="updateAttachmentData.FileName"/>
          </el-form-item>
          <el-form-item label="FilePath" :label-width="labelWidth">
            <el-input v-model="updateAttachmentData.FilePath"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="updateAttachmentFormVisible = false">Cancel</el-button>
          <el-button @click="UpdateAttachment">Submit</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="deleteAttachmentFormVisible"
        title="Delete Attachment"
        width="500"
        @close="ClearDeleteAttachmentForm"
      >
        <el-text>Are you confirm to delete the attachment?</el-text>
        <template #footer>
          <el-button @click="deleteAttachmentFormVisible = false">Cancel</el-button>
          <el-button @click="DeleteAttachment">Confirm</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="uploadAttachmentFormVisible"
        title="Upload Attachment"
        width="500"
        @close="ClearUploadAttachmentForm"
        >
        <el-form :model="uploadAttachmentData">
          <el-form-item label="FileName" :label-width="labelWidth">
            <el-input v-model="uploadAttachmentData.fileName"/>
          </el-form-item>
          <el-form-item label="File" :label-width="labelWidth">
            <el-upload
              style="width: 100%;"
              ref="upload"
              :limit="1"
              :auto-upload="false"
              :on-change="handleFileChange"
              :on-exceed="handleExceed"
            >
              <template #trigger>
                <el-button>Select</el-button>
              </template>
            </el-upload>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="uploadAttachmentFormVisible = false">Cancel</el-button>
          <el-button @click="UploadAttachment">Confirm</el-button>
        </template>
      </el-dialog>
    </el-main>
  </el-container>
</template>
<script>
import attachmentApi from "@/api/attachment.js"
import { ElMessage } from "element-plus";

export default {
  data() {
    return {
      labelWidth: 100,
      attachmentList: [],
      createAttachmentFormVisible: false,
      createAttachmentData: {
        "FileName": "",
        "FilePath": ""
      },
      updateAttachmentFormVisible: false,
      updateAttachmentData: {
        "ID": 0,
        "FileName": "",
        "FilePath": ""
      },
      deleteAttachmentFormVisible: false,
      deleteAttachmentData: {
        "ID": 0,
      },
      uploadAttachmentFormVisible: false,
      uploadAttachmentData: {
        "fileName": "",
        "file": null
      }
    }
  },
  methods: {
    GetAttachmentList() {
      attachmentApi.GetAttachmentList().then(res => {
        if (res.code === 0) {
          this.attachmentList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    CreateAttachment() {
      attachmentApi.CreateAttachment(this.createAttachmentData).then(res => {
        if (res.code === 0) {
          this.createAttachmentFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetAttachmentList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    ClearCreateAttachmentForm() {
      this.createAttachmentData = {
        "FileName": "",
        "FilePath": ""
      }
    },
    UpdateAttachment() {
      attachmentApi.UpdateAttachment(this.updateAttachmentData).then(res => {
        if (res.code === 0) {
          this.updateAttachmentFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetAttachmentList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenUpdateAttachmentForm(row) {
      this.updateAttachmentData = {
        "ID": row.ID,
        "FileName": row.FileName,
        "FilePath": row.FilePath
      }
      this.updateAttachmentFormVisible = true
    },
    ClearUpdateAttachmentForm() {
      this.updateAttachmentData = {
        "ID": 0,
        "FileName": "",
        "FilePath": ""
      }
    },
    DeleteAttachment() {
      attachmentApi.DeleteAttachment(this.deleteAttachmentData).then(res => {
        if (res.code === 0) {
          this.deleteAttachmentFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetAttachmentList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenDeleteAttachmentForm(row) {
      this.deleteAttachmentData = {
        "ID": row.ID,
      }
      this.deleteAttachmentFormVisible = true
    },
    ClearDeleteAttachmentForm() {
      this.deleteAttachmentData = {
        "ID": 0,
      }
    },
    UploadAttachment() {
      attachmentApi.UploadAttachment(this.uploadAttachmentData.fileName, this.uploadAttachmentData.file).then(res => {
        if (res.code === 0) {
          this.uploadAttachmentFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetAttachmentList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    ClearUploadAttachmentForm() {
      this.uploadAttachmentData = {
        "fileName": "",
        "file": null
      }
      this.$refs.upload.clearFiles();
    },
    DownloadAttachment(id) {
      attachmentApi.DownloadAttachment(id).then(ok => {
        if (!ok) {
          ElMessage.error("Download attachment fail")
        }
      }).catch(error => {
        console.log(error)
      })
    },
    handleFileChange(file) {
      if (file) {
        this.uploadAttachmentData.fileName = file.name;
        this.uploadAttachmentData.file = file.raw;
      } else {
        this.uploadAttachmentData.fileName = '';
        this.uploadAttachmentData.file = null;
      }
    },
    handleExceed(files) {
      const uploadInstance = this.$refs.upload;
      uploadInstance.clearFiles();
      const file = files[0];
      uploadInstance.handleStart(file);
    }
  },
  mounted() {
    this.GetAttachmentList()
  }
}
</script>
<style scoped>

</style>