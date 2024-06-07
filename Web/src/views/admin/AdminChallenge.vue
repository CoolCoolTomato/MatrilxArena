<template>
  <el-container>
    <el-header>
      <h2 style="color: var(--el-text-color-primary)">Challenge Manager</h2>
    </el-header>
    <el-main>
      <el-button style="margin: 10px" @click="createChallengeFormVisible = true">Add</el-button>
      <el-table
        :data="challengeList"
        table-layout="fixed"
        >
        <el-table-column prop="Title" label="Title"/>
        <el-table-column prop="Description" label="Description"/>
        <el-table-column prop="Image.RepoTags" label="Image"/>
        <el-table-column fixed="right" label="Operations" width="300px">
          <template #default=scope>
            <el-button
              @click="OpenUpdateChallengeForm(scope.row)"
              >
              Update
            </el-button>
            <el-button
              @click="OpenDeleteChallengeForm(scope.row)"
              >
              Delete
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-dialog
        v-model="createChallengeFormVisible"
        title="Create Challenge"
        width="500"
        @close="ClearCreateChallengeForm"
        >
        <el-form :model=createChallengeData>
          <el-form-item label="Title">
            <el-input v-model="createChallengeData.Title"/>
          </el-form-item>
          <el-form-item label="Description">
            <el-input v-model="createChallengeData.Description"/>
          </el-form-item>
          <el-form-item label="Image">
            <el-select
              v-model="createChallengeData.ImageID"
              filterable
              placeholder="Select"
              >
              <el-option
                v-for="image in imageList"
                :key="image.RepoTags"
                :label="image.RepoTags"
                :value="image.ID"
              >
              </el-option>
            </el-select>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="createChallengeFormVisible = false">Cancel</el-button>
          <el-button @click="CreateChallenge">Submit</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="updateChallengeFormVisible"
        title="Update Challenge"
        width="500"
        @close="ClearUpdateChallengeForm"
        >
        <el-form :model=updateChallengeData>
          <el-form-item label="Title">
            <el-input v-model="updateChallengeData.Title"/>
          </el-form-item>
          <el-form-item label="Description">
            <el-input v-model="updateChallengeData.Description"/>
          </el-form-item>
          <el-form-item label="Image">
            <el-select
              v-model="updateChallengeData.ImageID"
              filterable
              placeholder="Select"
              >
              <el-option
                v-for="image in imageList"
                :key="image.RepoTags"
                :label="image.RepoTags"
                :value="image.ID"
                >
              </el-option>
            </el-select>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="updateChallengeFormVisible = false">Cancel</el-button>
          <el-button @click="UpdateChallenge">Submit</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="deleteChallengeFormVisible"
        title="Delete Challenge"
        width="500"
        @close="ClearDeleteChallengeForm"
        >
        <el-text>Are you confirm to delete the challenge?</el-text>
        <template #footer>
        <el-button @click="deleteChallengeFormVisible = false">Cancel</el-button>
        <el-button @click="DeleteChallenge">Confirm</el-button>
        </template>
      </el-dialog>
    </el-main>
  </el-container>
</template>
<script>
import challengeApi from "@/api/challenge.js"
import imageApi from "@/api/image.js";
import { ElMessage } from "element-plus";

export default {
  data() {
    return {
      imageList: [],
      challengeList: [],
      createChallengeFormVisible: false,
      createChallengeData: {
        "ImageID": null,
        "Title": "",
        "Description": "",
      },
      updateChallengeFormVisible: false,
      updateChallengeData: {
        "ID": 0,
        "ImageID": null,
        "Title": "",
        "Description": "",
      },
      deleteChallengeFormVisible: false,
      deleteChallengeData: {
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
    GetChallengeList() {
      challengeApi.GetChallengeList().then(res => {
        if (res.code === 0) {
          this.challengeList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    CreateChallenge() {
      challengeApi.CreateChallenge(this.createChallengeData).then(res => {
        if (res.code === 0) {
          this.createChallengeFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetChallengeList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    ClearCreateChallengeForm() {
      this.createChallengeData = {
        "ImageID": null,
        "Title": "",
        "Description": "",
      }
    },
    UpdateChallenge() {
      challengeApi.UpdateChallenge(this.updateChallengeData).then(res => {
        if (res.code === 0) {
          this.updateChallengeFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetChallengeList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenUpdateChallengeForm(row) {
      this.updateChallengeData = {
        "ID": row.ID,
        "ImageID": row.ImageID,
        "Title": row.Title,
        "Description": row.Description,
      }
      this.updateChallengeFormVisible = true
    },
    ClearUpdateChallengeForm() {
      this.updateChallengeData = {
        "ID": 0,
        "ImageID": null,
        "Title": "",
        "Description": "",
      }
    },
    DeleteChallenge() {
      challengeApi.DeleteChallenge(this.deleteChallengeData).then(res => {
        if (res.code === 0) {
          this.deleteChallengeFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
            })
          this.GetChallengeList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenDeleteChallengeForm(row) {
      this.deleteChallengeData = {
        "ID": row.ID,
      }
      this.deleteChallengeFormVisible = true
    },
    ClearDeleteChallengeForm() {
      this.deleteChallengeData = {
        "ID": 0,
      }
    },
  },
  mounted() {
    this.GetImageList()
    this.GetChallengeList()
  }
}
</script>
