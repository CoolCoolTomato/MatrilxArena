<template>
  <el-container>
    <el-header>
      <div style="display:flex; align-items: center;">
        <h2 style="color: var(--el-text-color-primary);">{{ $t('AdminAttachment.AttachmentManager') }}</h2>
        <div style="flex-grow: 1;"></div>
        <div style="margin-right: 50px;">
          <el-button
            style="margin: 10px;"
            @click="createAttachmentFormVisible = true"
            type="primary"
          >
            {{ $t('AdminAttachment.Add') }}
          </el-button>
          <el-button
            style="margin: 10px;"
            @click="uploadAttachmentFormVisible = true"
          >
            {{ $t('AdminAttachment.Upload') }}
          </el-button>
        </div>
      </div>
    </el-header>
    <el-main>
      <el-table
        :data="attachmentList"
        table-layout="fixed"
        height="100%"
      >
        <el-table-column prop="FileName" :label="$t('AdminAttachment.FileName')"/>
        <el-table-column prop="FilePath" :label="$t('AdminAttachment.FilePath')"/>
        <el-table-column fixed="right" :label="$t('AdminAttachment.Operations')" width="320px">
          <template #default=scope>
            <el-button
              @click="OpenUpdateAttachmentForm(scope.row)"
            >
              {{ $t('AdminAttachment.Update') }}
            </el-button>
            <el-button
              @click="OpenDeleteAttachmentForm(scope.row)"
            >
              {{ $t('AdminAttachment.Delete') }}
            </el-button>
            <el-button
              @click="DownloadAttachment(scope.row)"
              >
              {{ $t('AdminAttachment.Download') }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-dialog
        v-model="createAttachmentFormVisible"
        :title="$t('AdminAttachment.CreateAttachment')"
        width="500"
        @close="ClearCreateAttachmentForm"
      >
        <el-form :model=createAttachmentData>
          <el-form-item :label="$t('AdminAttachment.FileName')" :label-width="labelWidth">
            <el-input v-model="createAttachmentData.FileName"/>
          </el-form-item>
          <el-form-item :label="$t('AdminAttachment.FilePath')" :label-width="labelWidth">
            <el-input v-model="createAttachmentData.FilePath"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button
            @click="CreateAttachment"
            type="primary"
            >
            {{ $t('AdminAttachment.Submit') }}
          </el-button>
          <el-button
            @click="createAttachmentFormVisible = false"
            >
            {{ $t('AdminAttachment.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="updateAttachmentFormVisible"
        :title="$t('AdminAttachment.UpdateAttachment')"
        width="500"
        @close="ClearUpdateAttachmentForm"
      >
        <el-form :model=updateAttachmentData>
          <el-form-item :label="$t('AdminAttachment.FileName')" :label-width="labelWidth">
            <el-input v-model="updateAttachmentData.FileName"/>
          </el-form-item>
          <el-form-item :label="$t('AdminAttachment.FilePath')" :label-width="labelWidth">
            <el-input v-model="updateAttachmentData.FilePath"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button
            @click="UpdateAttachment"
            type="primary"
            >
            {{ $t('AdminAttachment.Submit') }}
          </el-button>
          <el-button
            @click="updateAttachmentFormVisible = false"
            >
          {{ $t('AdminAttachment.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="deleteAttachmentFormVisible"
        :title="$t('AdminAttachment.DeleteAttachment')"
        width="500"
        @close="ClearDeleteAttachmentForm"
      >
        <el-text>{{ $t('AdminAttachment.AreYouConfirmToDeleteTheAttachment') }}</el-text>
        <template #footer>
          <el-button
            @click="DeleteAttachment"
            type="primary"
            >
            {{ $t('AdminAttachment.Confirm') }}
          </el-button>
          <el-button
            @click="deleteAttachmentFormVisible = false"
            >
            {{ $t('AdminAttachment.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="uploadAttachmentFormVisible"
        :title="$t('AdminAttachment.UploadAttachment')"
        width="500"
        @close="ClearUploadAttachmentForm"
        >
        <el-form :model="uploadAttachmentData">
          <el-form-item :label="$t('AdminAttachment.FileName')" :label-width="labelWidth">
            <el-input v-model="uploadAttachmentData.fileName"/>
          </el-form-item>
          <el-form-item :label="$t('AdminAttachment.File')" :label-width="labelWidth">
            <el-upload
              style="width: 100%;"
              ref="upload"
              :limit="1"
              :auto-upload="false"
              :on-change="handleFileChange"
              :on-exceed="handleExceed"
            >
              <template #trigger>
                <el-button size="small">{{ $t('AdminAttachment.SelectFile') }}</el-button>
              </template>
            </el-upload>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button
            @click="UploadAttachment"
            type="primary"
            >
            {{ $t('AdminAttachment.Confirm') }}
          </el-button>
          <el-button
            @click="uploadAttachmentFormVisible = false"
            >
            {{ $t('AdminAttachment.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
    </el-main>
  </el-container>
</template>
<script>
import attachmentApi from "@/api/attachment.js"
import { ElMessage } from "element-plus"
import { useI18n } from "vue-i18n"

export default {
  setup() {
    const { t } = useI18n()
    return { t }
  },
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
      this.$refs.upload.clearFiles()
    },
    DownloadAttachment(attachment) {
      try {
        new URL(attachment.FilePath)
        window.open(attachment.FilePath, '_blank')
      } catch (err) {
        attachmentApi.DownloadAttachment(attachment.ID).then(ok => {
          if (!ok) {
            ElMessage.error(this.t('AdminAttachment.DownloadAttachmentFail'))
          }
        }).catch(error => {
          console.log(error)
        })
      }
    },
    handleFileChange(file) {
      if (file) {
        this.uploadAttachmentData.fileName = file.name
        this.uploadAttachmentData.file = file.raw
      } else {
        this.uploadAttachmentData.fileName = ''
        this.uploadAttachmentData.file = null
      }
    },
    handleExceed(files) {
      const uploadInstance = this.$refs.upload
      uploadInstance.clearFiles()
      const file = files[0]
      uploadInstance.handleStart(file)
    }
  },
  mounted() {
    this.GetAttachmentList()
  }
}
</script>
