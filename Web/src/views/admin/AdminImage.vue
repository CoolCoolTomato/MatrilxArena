<template>
  <el-container>
    <el-header>
      <h2 style="color: var(--el-text-color-primary)">{{ t('AdminImage.ImageManager') }}</h2>
    </el-header>
    <el-main>
      <el-button
        style="margin: 10px"
        @click="createImageFormVisible = true"
        type="primary"
        >
        {{ t('AdminImage.Add') }}
      </el-button>
      <el-table
        :data="imageList"
        table-layout="fixed"
        >
        <el-table-column prop="Remark" :label="t('AdminImage.Remark')"/>
        <el-table-column prop="RepoTags" :label="t('AdminImage.RepoTags')"/>
        <el-table-column :label="t('AdminImage.Repository')">
          <template #default="scope">
            {{ formatRepository(scope.row.Repository) }}
          </template>
        </el-table-column>
        <el-table-column fixed="right" :label="t('AdminImage.Operations')" width="300px">
          <template #default=scope>
            <el-button
              @click="OpenUpdateImageForm(scope.row)"
              >
              {{ t('AdminImage.Update') }}
            </el-button>
            <el-button
              @click="OpenDeleteImageForm(scope.row)"
              >
              {{ t('AdminImage.Delete') }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-dialog
        v-model="createImageFormVisible"
        :title="t('AdminImage.CreateImage')"
        width="500"
        @close="ClearCreateImageForm"
        >
        <el-form :model=createImageData>
          <el-form-item :label="t('AdminImage.Remark')" :label-width="labelWidth">
            <el-input v-model="createImageData.Remark"/>
            </el-form-item>
          <el-form-item :label="t('AdminImage.RepoTags')" :label-width="labelWidth">
            <el-input v-model="createImageData.RepoTags"/>
          </el-form-item>
          <el-form-item :label="t('AdminImage.Repository')" :label-width="labelWidth">
            <el-input v-model="createImageData.Repository"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button
            @click="CreateImage"
            type="primary"
            >
            {{ t('AdminImage.Submit') }}
          </el-button>
          <el-button
            @click="createImageFormVisible = false"
            >
          {{ t('AdminImage.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="updateImageFormVisible"
        :title="t('AdminImage.UpdateImage')"
        width="500"
        @close="ClearUpdateImageForm"
        >
        <el-form :model=updateImageData>
          <el-form-item :label="t('AdminImage.Remark')" :label-width="labelWidth">
            <el-input v-model="updateImageData.Remark"/>
          </el-form-item>
          <el-form-item :label="t('AdminImage.RepoTags')" :label-width="labelWidth">
            <el-input v-model="updateImageData.RepoTags"/>
          </el-form-item>
          <el-form-item :label="t('AdminImage.Repository')" :label-width="labelWidth">
            <el-input v-model="updateImageData.Repository"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button
            @click="UpdateImage"
            type="primary"
            >
            {{ t('AdminImage.Submit') }}
          </el-button>
          <el-button
            @click="updateImageFormVisible = false"
            >
          {{ t('AdminImage.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="deleteImageFormVisible"
        :title="t('AdminImage.DeleteImage')"
        width="500"
        @close="ClearDeleteImageForm"
        >
        <el-text>{{ t('AdminImage.AreYouConfirmToDeleteTheImage') }}</el-text>
        <template #footer>
          <el-button
            @click="DeleteImage"
            type="primary"
            >
            {{ t('AdminImage.Confirm') }}
          </el-button>
          <el-button
            @click="deleteImageFormVisible = false"
            >
            {{ t('AdminImage.Cancel') }}
          </el-button>
        </template>
      </el-dialog>
    </el-main>
  </el-container>
</template>
<script>
import imageApi from "@/api/image.js"
import { ElMessage } from "element-plus";
import { useI18n } from "vue-i18n";

export default {
  setup() {
    const { t } = useI18n()
    return { t }
  },
  data() {
    return {
      labelWidth: 100,
      imageList: [],
      createImageFormVisible: false,
      createImageData: {
        "Remark": "",
        "RepoTags": "",
        "Repository": "",
      },
      updateImageFormVisible: false,
      updateImageData: {
        "ID": 0,
        "Remark": "",
        "RepoTags": "",
        "Repository": "",
      },
      deleteImageFormVisible: false,
      deleteImageData: {
        "ID": 0,
      }
    }
  },
  methods: {
    GetImageList() {
      imageApi.GetImageList().then(res => {
        if (res.code === 0) {
          this.imageList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    CreateImage() {
      imageApi.CreateImage(this.createImageData).then(res => {
        if (res.code === 0) {
          this.createImageFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetImageList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    ClearCreateImageForm() {
      this.createImageData = {
        "Remark": "",
        "RepoTags": "",
        "Repository": "",
      }
    },
    UpdateImage() {
      imageApi.UpdateImage(this.updateImageData).then(res => {
        if (res.code === 0) {
          this.updateImageFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetImageList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenUpdateImageForm(row) {
      this.updateImageData = {
        "ID": row.ID,
        "Remark": row.Remark,
        "RepoTags": row.RepoTags,
        "Repository": row.Repository,
      }
      this.updateImageFormVisible = true
    },
    ClearUpdateImageForm() {
      this.updateImageData = {
        "ID": 0,
        "Remark": "",
        "RepoTags": "",
        "Repository": "",
      }
    },
    DeleteImage() {
      imageApi.DeleteImage(this.deleteImageData).then(res => {
        if (res.code === 0) {
          this.deleteImageFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetImageList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenDeleteImageForm(row) {
      this.deleteImageData = {
        "ID": row.ID,
      }
      this.deleteImageFormVisible = true
    },
    ClearDeleteImageForm() {
      this.deleteImageData = {
        "ID": 0,
      }
    },
    formatRepository(repository) {
      return repository === '' ? this.t('AdminImage.NotSpecified') : repository;
    },
  },
  mounted() {
    this.GetImageList()
  }
}
</script>
