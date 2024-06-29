<template>
  <el-container>
    <el-header>
      <h2 style="color: var(--el-text-color-primary)">ChallengeClass Manager</h2>
    </el-header>
    <el-main>
      <el-button style="margin: 10px" @click="createChallengeClassFormVisible = true">Add</el-button>
      <el-table
        :data="challengeClassList"
        table-layout="fixed"
      >
        <el-table-column prop="Name" label="Name"/>
        <el-table-column fixed="right" label="Operations" width="300px">
          <template #default=scope>
            <el-button
              @click="OpenUpdateChallengeClassForm(scope.row)"
            >
              Update
            </el-button>
            <el-button
              @click="OpenDeleteChallengeClassForm(scope.row)"
            >
              Delete
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-dialog
        v-model="createChallengeClassFormVisible"
        title="Create ChallengeClass"
        width="500"
        @close="ClearCreateChallengeClassForm"
      >
        <el-form :model=createChallengeClassData>
          <el-form-item label="Name" :label-width="labelWidth">
            <el-input v-model="createChallengeClassData.Name"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="createChallengeClassFormVisible = false">Cancel</el-button>
          <el-button @click="CreateChallengeClass">Submit</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="updateChallengeClassFormVisible"
        title="Update ChallengeClass"
        width="500"
        @close="ClearUpdateChallengeClassForm"
      >
        <el-form :model=updateChallengeClassData>
          <el-form-item label="Name" :label-width="labelWidth">
            <el-input v-model="updateChallengeClassData.Remark"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="updateChallengeClassFormVisible = false">Cancel</el-button>
          <el-button @click="UpdateChallengeClass">Submit</el-button>
        </template>
      </el-dialog>
      <el-dialog
        v-model="deleteChallengeClassFormVisible"
        title="Delete ChallengeClass"
        width="500"
        @close="ClearDeleteChallengeClassForm"
      >
        <el-text>Are you confirm to delete the challengeClass?</el-text>
        <template #footer>
          <el-button @click="deleteChallengeClassFormVisible = false">Cancel</el-button>
          <el-button @click="DeleteChallengeClass">Confirm</el-button>
        </template>
      </el-dialog>
    </el-main>
  </el-container>
</template>
<script>
import challengeClassApi from "@/api/challengeClass.js"
import { ElMessage } from "element-plus";

export default {
  data() {
    return {
      labelWidth: 100,
      challengeClassList: [],
      createChallengeClassFormVisible: false,
      createChallengeClassData: {
        "Name": "",
      },
      updateChallengeClassFormVisible: false,
      updateChallengeClassData: {
        "Name": "",
      },
      deleteChallengeClassFormVisible: false,
      deleteChallengeClassData: {
        "ID": 0,
      }
    }
  },
  methods: {
    GetChallengeClassList() {
      challengeClassApi.GetChallengeClassList().then(res => {
        if (res.code === 0) {
          this.challengeClassList = res.data
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    CreateChallengeClass() {
      challengeClassApi.CreateChallengeClass(this.createChallengeClassData).then(res => {
        if (res.code === 0) {
          this.createChallengeClassFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetChallengeClassList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    ClearCreateChallengeClassForm() {
      this.createChallengeClassData = {
        "Name": "",
      }
    },
    UpdateChallengeClass() {
      challengeClassApi.UpdateChallengeClass(this.updateChallengeClassData).then(res => {
        if (res.code === 0) {
          this.updateChallengeClassFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetChallengeClassList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenUpdateChallengeClassForm(row) {
      this.updateChallengeClassData = {
        "ID": row.ID,
        "Name": row.Name,
      }
      this.updateChallengeClassFormVisible = true
    },
    ClearUpdateChallengeClassForm() {
      this.updateChallengeClassData = {
        "ID": 0,
        "Name": "",
      }
    },
    DeleteChallengeClass() {
      challengeClassApi.DeleteChallengeClass(this.deleteChallengeClassData).then(res => {
        if (res.code === 0) {
          this.deleteChallengeClassFormVisible = false
          ElMessage({
            message: res.msg,
            type: 'success',
          })
          this.GetChallengeClassList()
        } else {
          ElMessage.error(res.msg)
        }
      }).catch(error => {
        console.log(error)
      })
    },
    OpenDeleteChallengeClassForm(row) {
      this.deleteChallengeClassData = {
        "ID": row.ID,
      }
      this.deleteChallengeClassFormVisible = true
    },
    ClearDeleteChallengeClassForm() {
      this.deleteChallengeClassData = {
        "ID": 0,
      }
    },
  },
  mounted() {
    this.GetChallengeClassList()
  }
}
</script>
