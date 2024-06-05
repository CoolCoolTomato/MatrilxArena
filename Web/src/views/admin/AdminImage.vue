<template>
  <el-container>
    <el-header>
      <h2 style="color: var(--el-text-color-primary)">Image Manager</h2>
    </el-header>
    <el-main>
      <el-button style="margin: 10px" @click="createImageFormVisible = true">Add</el-button>
      <el-table
        :data="imageList"
        table-layout="fixed"
        >
        <el-table-column prop="RepoTags" label="RepoTags"/>
        <el-table-column label="Repository">
          <template #default="scope">
            {{ formatRepository(scope.row.Repository) }}
          </template>
        </el-table-column>
        <el-table-column fixed="right" label="Operations" width="300px">
          <template #default=scope>
            <el-button
              @click="OpenUpdateImageForm(scope.row)"
              >
              Update
            </el-button>
            <el-button
              @click="OpenDeleteImageForm(scope.row)"
              >
              Delete
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-dialog
        v-model="createImageFormVisible"
        title="Create Image"
        width="500"
        @close="ClearCreateImageForm"
        >
        <el-form :model=createImageData>
          <el-form-item label="RepoTags">
            <el-input v-model="createImageData.RepoTags"/>
          </el-form-item>
          <el-form-item label="Repository">
            <el-input v-model="createImageData.Repository"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="createImageFormVisible = false">Cancel</el-button>
          <el-button @click="CreateImage">Submit</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="updateImageFormVisible"
        title="Update Image"
        width="500"
        @close="ClearUpdateImageForm"
        >
        <el-form :model=updateImageData>
          <el-form-item label="RepoTags">
            <el-input v-model="updateImageData.RepoTags"/>
          </el-form-item>
          <el-form-item label="Repository">
            <el-input v-model="updateImageData.Repository"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="updateImageFormVisible = false">Cancel</el-button>
          <el-button @click="UpdateImage">Submit</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="deleteImageFormVisible"
        title="Delete Image"
        width="500"
        @close="ClearDeleteImageForm"
        >
        <el-text>Are you confirm to delete the image?</el-text>
        <template #footer>
        <el-button @click="deleteImageFormVisible = false">Cancel</el-button>
        <el-button @click="DeleteImage">Confirm</el-button>
        </template>
      </el-dialog>
    </el-main>
  </el-container>
</template>
<script>
import imageApi from "@/api/image.js"
import { ElMessage } from "element-plus";

export default {
  data() {
    return {
      imageList: [],
      createImageFormVisible: false,
      createImageData: {
        "RepoTags": "",
        "Repository": "",
      },
      updateImageFormVisible: false,
      updateImageData: {
        "ID": 0,
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
        "RepoTags": row.RepoTags,
        "Repository": row.Repository,
      }
      this.updateImageFormVisible = true
    },
    ClearUpdateImageForm() {
      this.updateImageData = {
        "ID": 0,
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
      return repository === '' ? 'Not Specified' : repository;
    },
  },
  mounted() {
    this.GetImageList()
  }
}
</script>
